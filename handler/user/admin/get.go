package admin

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Get gets a user account.
//
// @Summary Get a user information
// @Description Get a user account specified by user ID
// @Tags user/admin
// @Produce  json
// @Param id path uint64 true "the ID of the specified user to update"
// @Success 200 {object} user.SwaggerGetResponse "{"code":0,"message":"OK","data":{"username":"傅秀英","password":"$2a$10$5pLrLpEQ1HAD2Hcm3Bnud.Shhmf5bTaf1yTWYloot0i5nvn1Td4hq","role":"general"}}"
// @Router /user/admin/{id} [get]
// @Security ApiKeyAuth
func Get(c *gin.Context) {
	log.GetUserCalled(c)

	userId, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, nil, err)
		return
	}
	// Get the user by the `username` from the database.
	user, err := model.GetUserByID(userId)
	// if user is not found or deleted, send error
	if err != nil {
		SendResponse(c, berror.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
