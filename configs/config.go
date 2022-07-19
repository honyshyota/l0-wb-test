package config

import (
	"errors"
	"fmt"

	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Config struct {
	App struct {
		Port string
		Host string
	}
	DB struct {
		DB *sqlx.DB
	}
	NatsStreaming struct {
		ClientID string
		Cluster  string
		Path     string
		Host     string
		Port     string
		Time     time.Duration
	}
}

func NewConfig() (*Config, error) {
	var once sync.Once
	var config *Config

	once.Do(func() {
		err := godotenv.Load("configs/conf.env")
		if err != nil {
			logrus.Fatal("Cannot to load env", err)
		}

		timeSub := os.Getenv("TIME_SUB")
		hostDB := os.Getenv("HOST_DB")
		nameDB := os.Getenv("NAME_DB")
		user := os.Getenv("USER_DB")
		password := os.Getenv("PASSWORD")
		portDB := os.Getenv("PORT_DB")
		sqlConn, err := connToDB(portDB, hostDB, user, password, nameDB)
		if err != nil {
			logrus.Fatal("Unable to connect DB ", err)
		}
		config = &Config{
			App: struct {
				Port string
				Host string
			}{
				Port: os.Getenv("PORT_APP"),
				Host: os.Getenv("HOST_APP"),
			},
			DB: struct{ DB *sqlx.DB }{
				DB: sqlConn,
			},
			NatsStreaming: struct {
				ClientID string
				Cluster  string
				Path     string
				Host     string
				Port     string
				Time     time.Duration
			}{
				ClientID: os.Getenv("CLIENT_ID_NS"),
				Cluster:  os.Getenv("CLUSTER_NS"),
				Path:     os.Getenv("PATH_NS"),
				Port:     os.Getenv("PORT_NS"),
				Host:     os.Getenv("HOST_NS"),
				Time:     time.Since(getTime(timeSub)),
			},
		}
	})

	return config, nil
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

	if err := runPgMigrations(psqlconnect); err != nil {
		return nil, fmt.Errorf("runPgMigrations failed: %w", err)
	}
	logrus.Println("migration done")

	return db, nil
}

func getTime(timeSub string) time.Time {
	timeS := strings.Split(timeSub, ":")
	year, err := strconv.Atoi(timeS[0])
	if err != nil {
		return time.Now()
	}

	month, err := strconv.Atoi(timeS[1])
	if err != nil {
		return time.Now()
	}

	day, err := strconv.Atoi(timeS[2])
	if err != nil {
		return time.Now()
	}

	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

}

func runPgMigrations(pgURL string) error {
	MigrationsPath := os.Getenv("PG_MIGRATIONS_PATH")

	if MigrationsPath == "" {
		return nil
	}

	if pgURL == "" {
		return errors.New("no PgURL provided")
	}

	m, err := migrate.New(
		MigrationsPath,
		pgURL,
	)
	if err != nil {

		return err
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}

	return nil
}
