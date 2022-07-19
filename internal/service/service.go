package service

import (
	"encoding/json"

	config "github.com/honyshyota/l0-wb-test/configs"
	"github.com/honyshyota/l0-wb-test/internal/model"
	"github.com/honyshyota/l0-wb-test/internal/store/cache"
	"github.com/honyshyota/l0-wb-test/internal/store/db"
	"github.com/sirupsen/logrus"
)

type service struct {
	cache cache.Cache
	repo  db.Repo
}

func NewService(config *config.Config) Service {
	cache := cache.NewCache()
	logrus.Println("cache is ready")

	db := db.NewDB(config)
	logrus.Println("DB is ready")

	return &service{
		cache: cache,
		repo:  db,
	}
}

func (s *service) FindById(id int) *model.Order {
	return s.cache.Get(id)
}

func (s *service) FindByIdBadMessage(id int) []byte {
	return s.cache.GetBadCache(id)
}

func (s *service) Set(message []byte) error {
	order := &model.Order{}
	var id int

	err := json.Unmarshal(message, &order)
	if err != nil {
		
		return err
	}

	if order.OrderUID == "" {
		id, err = s.repo.CreateBadMessage(message)
		if err != nil {
			return err
		}

		s.cache.SetBadCache(id, message)
	} else {
		id, err = s.repo.Create(order)
		if err != nil {
			return err
		}

		s.cache.Set(id, order)
	}

	return nil
}

func (s *service) GetAll() error {
	rows, rowsBadMessages, err := s.repo.FindAll()
	if err != nil {
		return err
	}

	defer rows.Close()
	defer rowsBadMessages.Close()

	var id int
	var data []byte
	var orders *model.Order

	for rows.Next() {
		err = rows.Scan(&id, &data)
		if err != nil {
			return err
		}

		err = json.Unmarshal(data, &orders)
		if err != nil {
			return err
		}

		s.cache.Set(id, orders)
	}

	for rowsBadMessages.Next() {
		err = rowsBadMessages.Scan(&id, &data)
		if err != nil {
			return err
		}

		s.cache.SetBadCache(id, data)

	}

	return nil
}
