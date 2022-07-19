package db

import (
	"github.com/honyshyota/l0-wb-test/internal/model"
	"github.com/jmoiron/sqlx"
)

type Repo interface {
	Create(*model.Order) (int, error)
	FindAll() (*sqlx.Rows, *sqlx.Rows, error)
	CreateBadMessage([]byte) (int, error)
}
