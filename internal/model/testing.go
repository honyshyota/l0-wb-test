package model

import (
	"testing"
	"time"
)

func TestOrder(t *testing.T) *Order {
	var Item []*Item
	Item = append(Item, TestItem(t))
	return &Order{
		OrderUID:          "b563feb7b2b84b6test",
		TrackNum:          "WBILMTESTTRACK",
		Entry:             "WBIL",
		Delivery:          TestDelivery(t),
		Payment:           TestPayment(t),
		Items:             Item,
		Locale:            "en",
		InternalSignature: "",
		CustomerID:        "test",
		DeliveryService:   "meest",
		ShardKey:          "9",
		SmID:              99,
		DateCreated:       time.Now(),
		OofShard:          "1",
	}
}

func TestDelivery(t *testing.T) *Delivery {
	return &Delivery{
		Name:    "Test Testov",
		Phone:   "+9720000000",
		Zip:     "2639809",
		City:    "Kiryat Mozkin",
		Address: "Ploshad Mira 15",
		Region:  "Kraiot",
		Email:   "test@gmail.com",
	}
}

func TestPayment(t *testing.T) *Payment {
	return &Payment{
		Transaction:  "b563feb7b2b84b6test",
		RequestID:    "",
		Currency:     "USD",
		Provider:     "wbpay",
		Amount:       1817,
		PaymentDT:    1637907727,
		Bank:         "alpha",
		DeliveryCost: 1500,
		GoodsTotal:   317,
		CustomFee:    0,
	}
}

func TestItem(t *testing.T) *Item {
	return &Item{
		ChrtID:      9934930,
		TrackNumber: "WBILMTESTTRACK",
		Price:       453,
		Rid:         "ab4219087a764ae0btest",
		Name:        "Mascaras",
		Sale:        30,
		Size:        "0",
		TotalPrice:  317,
		NmID:        2389212,
		Brand:       "Vivienne Sabo",
		Status:      202,
	}
}
