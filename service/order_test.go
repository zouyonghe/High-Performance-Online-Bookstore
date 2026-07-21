package service

import (
	"testing"
	"time"

	"High-Performance-Online-Bookstore/model"
)

func TestBuildOrderInfosIncludesUserID(t *testing.T) {
	orders := []*model.Order{{
		Base:       model.Base{ID: 7, CreatedAt: time.Date(2026, 7, 22, 8, 0, 0, 0, time.UTC)},
		UserID:     99,
		Books:      []model.OrderBook{{BookID: 1, UnitPrice: 12.5, Number: 2}},
		OrderPrice: 25,
		Status:     "open",
	}}

	infos := buildOrderInfos(orders)
	if len(infos) != 1 {
		t.Fatalf("len = %d, want 1", len(infos))
	}
	if infos[0].UserID != 99 {
		t.Fatalf("userID = %d, want 99", infos[0].UserID)
	}
	if infos[0].OrderID != 7 {
		t.Fatalf("orderID = %d, want 7", infos[0].OrderID)
	}
}
