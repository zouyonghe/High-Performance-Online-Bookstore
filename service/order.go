package service

import (
	. "High-Performance-Online-Bookstore/database"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/constvar"
	"sync"
)

func ListOrderInfo(userID uint64, pageNum int, pageSize int) ([]*model.OrderInfo, error) {
	infos := make([]*model.OrderInfo, 0)
	orders, err := ListOrder(userID, pageNum, pageSize)
	if err != nil {
		return nil, err
	}
	var ids []uint64
	for _, o := range orders {
		ids = append(ids, o.ID)
	}
	wg := sync.WaitGroup{}
	orderList := model.OrderList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint64]*model.OrderInfo, len(orders)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	for _, o := range orders {
		wg.Add(1)
		o := o
		go func(b *model.Order) {
			defer wg.Done()

			orderList.Lock.Lock()
			defer orderList.Lock.Unlock()

			orderList.IdMap[o.ID] = &model.OrderInfo{
				OrderID:    o.ID,
				Books:      o.Books,
				OrderPrice: o.OrderPrice,
				CreatedAt:  o.CreatedAt.Format("2006-01-02 15:04:05"),
				IsApproved: o.IsApproved,
			}
		}(o)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, err
	}
	for _, id := range ids {
		infos = append(infos, orderList.IdMap[id])
	}

	return infos, nil
}
func ListOrder(userID uint64, pageNum int, pageSize int) ([]*model.Order, error) {
	if pageSize <= 0 {
		pageSize = constvar.DefaultPageSize
	}
	var orders []*model.Order

	if pageNum <= 0 {
		pageNum = 1
	}
	offset := (pageNum - 1) * pageSize
	if err := DB.Self.Offset(offset).Limit(pageSize).Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}
	for _, order := range orders {
		var ob []model.OrderBook
		if err := DB.Self.Where("order_id = ?", order.ID).Find(&ob).Error; err != nil {
			return nil, err
		}
		order.Books = append(order.Books, ob...)
	}

	return orders, nil
}