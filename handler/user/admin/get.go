package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
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
	zap.L().Info("user get function called.", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	userId, _ := strconv.Atoi(c.Param("id"))
	// Get the user by the `username` from the database.
	user, err := model.GetUserByID(uint64(userId))
	// if user is not found or deleted, send error
	if err != nil {
		SendResponse(c, berror.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
