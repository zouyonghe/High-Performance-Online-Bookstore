package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// List handler returns a list of orders
func List(c *gin.Context) {
	log.ListOrderCalled(c)

	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBindRequest, nil)
		return
	}

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}

	// get orders
	orders, err := service.ListOrderInfo(userID, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListOrder(err)
		SendResponse(c, berror.InternalServerError, nil)
		return
	}
	SendResponse(c, nil, ListResponse{
		TotalCount: len(orders),
		OrderList:  orders,
	})
}
