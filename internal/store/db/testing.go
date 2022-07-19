package db

import (
	"fmt"
	"testing"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/jmoiron/sqlx"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func TestDB(t *testing.T) *config.Config {
	t.Helper()

	var conf *config.Config

	connToDB, _ := connToDB(":5432", "localhost", "postgres", "postgres", "wb_db_test")
	conf = &config.Config{
		DB: struct{ DB *sqlx.DB }{
			DB: connToDB,
		},
	}

	return conf
}

func connToDB(portDB, hostDB, user, password, nameDB string) (*sqlx.DB, error) {
	psqlconnect := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable", user, password, hostDB, portDB, nameDB)

	db, err := sqlx.Connect("pgx", psqlconnect)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
