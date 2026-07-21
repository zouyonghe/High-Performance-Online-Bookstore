package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// DealByID deals with one order by ID.
//
// @Summary Deal with one order by ID
// @Description Deal with one order by ID for manager roles
// @Tags order
// @Accept json
// @Produce json
// @Param id path uint64 true "the ID of the order"
// @Param order body order.DealOrderRequest true "order ID and operation"
// @Success 200 {object} order.SwaggerDealOrderResponse
// @Router /order/{id} [post]
// @Security ApiKeyAuth
func DealByID(c *gin.Context) {
	log.DealOrderCalled(c)

	orderID, err := service.GetIDByParam(c)
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

	o, err := model.GetOrder(orderID)
	if err != nil {
		log.ErrGetOrder(err)
		SendError(c, err)
		return
	}

	if err = o.DealWith(r.Operation); err != nil {
		log.ErrDealOrder(err)
		SendError(c, err)
		return
	}

	SendResponse(c, nil, DealOrderResponse{OrderID: orderID})
}
