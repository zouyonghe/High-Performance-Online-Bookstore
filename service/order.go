package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/pkg/constvar"
)

// ListOrderInfo lists the open and accepted orders of a user.
func ListOrderInfo(userID uint64, pageNum int, pageSize int) ([]*model.OrderInfo, error) {
	orders, err := ListOrder(userID, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return buildOrderInfos(orders), nil
}

// ListOrder lists the open and accepted orders of a user,
// loading the books of each order.
func ListOrder(userID uint64, pageNum int, pageSize int) ([]*model.Order, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	orders := make([]*model.Order, 0)
	if err := DB.Self.Offset(offset).Limit(pageSize).
		Where("user_id = ? and status IN ?", userID, []string{"open", "accept"}).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	if err := loadOrderBooks(orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// ListAcceptOrderInfo lists all accepted orders.
func ListAcceptOrderInfo(pageNum int, pageSize int) ([]*model.OrderInfo, error) {
	orders, err := ListAcceptOrder(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return buildOrderInfos(orders), nil
}

// ListAllOrderInfo lists all orders.
func ListAllOrderInfo(pageNum int, pageSize int) ([]*model.OrderInfo, error) {
	orders, err := ListAllOrder(pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	return buildOrderInfos(orders), nil
}

// ListAcceptOrder lists all accepted orders,
// loading the books of each order.
func ListAcceptOrder(pageNum int, pageSize int) ([]*model.Order, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	orders := make([]*model.Order, 0)
	if err := DB.Self.Offset(offset).Limit(pageSize).
		Where("status = ?", "accept").
		Find(&orders).Error; err != nil {
		return nil, err
	}
	if err := loadOrderBooks(orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// ListAllOrder lists all orders except deleted ones,
// loading the books of each order.
func ListAllOrder(pageNum int, pageSize int) ([]*model.Order, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize

	orders := make([]*model.Order, 0)
	if err := DB.Self.Offset(offset).Limit(pageSize).
		Find(&orders).Error; err != nil {
		return nil, err
	}
	if err := loadOrderBooks(orders); err != nil {
		return nil, err
	}
	return orders, nil
}

// List lists orders based on the role of the requester.
func List(role string, userID uint64, pageNum int, pageSize int) ([]*model.OrderInfo, error) {
	switch role {
	case "admin", "seller":
		return ListAllOrderInfo(pageNum, pageSize)
	case "general":
		return ListOrderInfo(userID, pageNum, pageSize)
	}
	return nil, berror.ErrIdentifyRole
}

// loadOrderBooks loads the books of each order in the list.
func loadOrderBooks(orders []*model.Order) error {
	for _, order := range orders {
		var ob []model.OrderBook
		if err := DB.Self.Where("order_id = ?", order.ID).Find(&ob).Error; err != nil {
			return err
		}
		order.Books = append(order.Books, ob...)
	}
	return nil
}

// buildOrderInfos converts order models to order information.
func buildOrderInfos(orders []*model.Order) []*model.OrderInfo {
	infos := make([]*model.OrderInfo, 0, len(orders))
	for _, o := range orders {
		infos = append(infos, &model.OrderInfo{
			OrderID:    o.ID,
			UserID:     o.UserID,
			Books:      o.Books,
			OrderPrice: o.OrderPrice,
			CreatedAt:  o.CreatedAt.Format("2006-01-02 15:04:05"),
			Status:     o.Status,
		})
	}
	return infos
}
