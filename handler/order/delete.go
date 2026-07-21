package order

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Delete deletes one order.
//
// @Summary Delete one order
// @Description Delete one order by ID for manager roles
// @Tags order
// @Produce json
// @Param id path uint64 true "the ID of the order"
// @Success 200 {object} order.SwaggerDealOrderResponse
// @Router /order/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	orderID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	o, err := model.GetOrder(orderID)
	if err != nil {
		log.ErrGetOrder(err)
		SendError(c, err)
		return
	}

	if err = o.DeleteOrder(); err != nil {
		log.ErrDeleteOrder(err)
		SendError(c, err)
		return
	}

	SendResponse(c, nil, DealOrderResponse{OrderID: orderID})
}
