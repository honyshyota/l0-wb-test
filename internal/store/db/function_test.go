package db_test

import (
	"io/ioutil"
	"testing"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/model"
	"github.com/honyshyota/l0-wb-test/internal/store/db"
	"github.com/stretchr/testify/assert"
)

func TestDb_Create(t *testing.T) {
	config, _ := config.NewConfig()

	s := db.NewDB(config)
	m := model.TestOrder(t)
	id, err := s.Create(m)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	config.DB.DB.QueryRowx("TRUNCATE orders CASCADE")
}

func TestDb_CreateBadMessage(t *testing.T) {
	config, _ := config.NewConfig()

	s := db.NewDB(config)
	model, err := ioutil.ReadFile("test_models/1.json")
	assert.NoError(t, err)
	id, err := s.CreateBadMessage(model)
	assert.NoError(t, err)
	assert.NotNil(t, id)
	config.DB.DB.QueryRowx("TRUNCATE bad_messages CASCADE")
}

func TestDb_FindAll(t *testing.T) {
	config, _ := config.NewConfig()

	s := db.NewDB(config)
	m1, m2, err := s.FindAll()
	assert.NotNil(t, m1)
	assert.NotNil(t, m2)
	assert.NoError(t, err)
}
