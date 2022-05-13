package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

func Deal(c *gin.Context) {
	log.DealOrderCalled(c)

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	var r DealOrderRequest
	if err = c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}

	o, err := model.GetOrder(r.OrderID)
	if err != nil {
		log.ErrGetOrder(err)
		SendError(c, err)
		return
	}
	err = o.CheckOwner(userID)
	if err != nil {
		log.ErrCheckOrderOwner(err)
		SendError(c, err)
		return
	}

	if err = o.DealWith(r.Operation); err != nil {
		log.ErrDealOrder(err)
		SendError(c, err)
		return
	}

	rsp := DealOrderResponse{
		OrderID: r.OrderID,
	}
	SendResponse(c, nil, rsp)
}
