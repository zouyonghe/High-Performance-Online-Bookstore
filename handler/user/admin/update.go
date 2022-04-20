package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// Update updates a user account.
//
// @Summary Update a user account
// @Description Update a user account specified by user ID
// @Tags user/admin
// @Accept  json
// @Produce  json
// @Param id path uint64 true "the ID of the specified user to update"
// @Param user body user.UpdateRequest true "user information include username and password"
// @Success 200 {object} user.SwaggerUpdateResponse "{"code":0,"message":"OK","data":{"userId":5}}"
// @Router /user/admin/{id} [put]
// @Security ApiKeyAuth
func Update(c *gin.Context) {
	zap.L().Info("update function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	var u model.UserModel
	// Binding the user data.
	if err := c.Bind(&u); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	var role string
	if userId != 0 {
		role = "general"
	} else {
		role = "admin"
	}

	u.ID = uint64(userId)
	u.Role = role

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, berror.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.Update(); err != nil {
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}
	rsp := user.UpdateResponse{
		UserId:   u.ID,
		Username: u.Username,
	}
	SendResponse(c, nil, rsp)

}
