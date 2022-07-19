package cache

import "github.com/honyshyota/l0-wb-test/internal/model"

type Cache interface {
	Set(int, *model.Order)
	Get(int) *model.Order
	SetBadCache(int, []byte)
	GetBadCache(int) []byte
}
