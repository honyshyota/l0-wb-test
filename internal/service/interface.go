package service

import "github.com/honyshyota/l0-wb-test/internal/model"

type Service interface {
	FindById(int) *model.Order
	FindByIdBadMessage(int) []byte
	Set([]byte) error
	GetAll() error
}
