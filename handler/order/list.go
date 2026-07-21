package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// List returns a list of orders.
//
// @Summary List orders
// @Description List orders: general users see their own orders, sellers and admins see all orders
// @Tags order
// @Produce json
// @Param pageNum query int false "page number, default 1"
// @Param pageSize query int false "page size, default 10"
// @Success 200 {object} order.SwaggerListResponse
// @Router /order [get]
// @Security ApiKeyAuth
func List(c *gin.Context) {
	log.ListOrderCalled(c)

	var r ListRequest
	if err := c.ShouldBindQuery(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}
	//get role from token
	role, err := service.GetRoleByToken(c)
	if err != nil {
		log.ErrGetRole(err)
		SendError(c, err)
		return
	}

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	// get orders
	orders, err := service.List(role, userID, r.PageNum, r.PageSize)
	if err != nil {
		log.ErrListOrder(err)
		SendError(c, err)
		return
	}
	SendResponse(c, nil, ListResponse{
		TotalCount: len(orders),
		OrderList:  orders,
	})
}
