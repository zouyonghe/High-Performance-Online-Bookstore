package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
)

func Deal(c *gin.Context) {
	log.DealOrderCalled(c)

	var r DealOrderRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}

	o, err := model.GetOrder(r.OrderID)
	if err != nil {
		log.ErrGetOrder(err)
		SendResponse(c, berror.ErrGetOrder, nil)
		return
	}

	if err = o.DealWith(r.Operation); err != nil {
		log.ErrDealOrder(err)
		SendResponse(c, err, nil)
		return
	}

	rsp := DealOrderResponse{
		OrderID: r.OrderID,
	}
	SendResponse(c, nil, rsp)
}
