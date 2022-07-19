package db

import (
	"encoding/json"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/model"
	"github.com/jmoiron/sqlx"
)

type repo struct {
	db *sqlx.DB
}

func NewDB(config *config.Config) Repo {
	return &repo{
		db: config.DB.DB,
	}
}

func (r *repo) Create(value *model.Order) (int, error) {
	var id int

	data, err := json.Marshal(value)
	if err != nil {
		return 0, err
	}

	insert := "INSERT INTO orders(data) VALUES($1::jsonb) RETURNING id"

	err = r.db.QueryRowx(insert, data).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) CreateBadMessage(value []byte) (int, error) {
	var id int

	insert := "INSERT INTO bad_messages(data) VALUES($1::jsonb) RETURNING id"

	err := r.db.QueryRowx(insert, value).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repo) FindAll() (*sqlx.Rows, *sqlx.Rows, error) {
	insert := "SELECT * FROM orders"

	insertBadMessages := "SELECT * FROM bad_messages"

	rows, err := r.db.Queryx(insert)
	if err != nil {
		return nil, nil, err
	}

	rowsBadMessages, err := r.db.Queryx(insertBadMessages)
	if err != nil {
		return nil, nil, err
	}

	return rows, rowsBadMessages, nil
}
