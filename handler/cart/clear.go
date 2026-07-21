package cart

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Clear clears the cart.
//
// @Summary Clear the cart
// @Description Remove all books from the current user's cart
// @Tags cart
// @Produce json
// @Success 200 {object} cart.SwaggerClearCartResponse
// @Router /cart/all [delete]
// @Security ApiKeyAuth
func Clear(c *gin.Context) {
	log.ClearCartCalled(c)

	// get user id
	userID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	// get cart
	cart, err := model.GetCart(userID)
	if err != nil {
		log.ErrGetCart(err)
		SendError(c, err)
		return
	}

	// clear cart
	err = cart.ClearCart()
	if err != nil {
		log.ErrClearCart(err)
		SendError(c, err)
		return
	}
	rsp := ClearCartResponse{
		Message: "cart cleared",
	}
	SendResponse(c, nil, rsp)
}
