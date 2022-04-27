package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
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
	log.UpdateUserCalled(c)
	// Get the user id from the url parameter.
	userId, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, nil, err)
		return
	}

	var u model.UserModel
	// Binding the user data.
	if err := c.ShouldBindJSON(&u); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	u.ID = userId

	m, err := model.GetUserByID(u.ID)
	if err != nil {
		log.ErrGetUser(err)
	}
	if m.Role == "admin" && u.Role != "admin" {
		zap.L().Error("admin user can't change role")
		SendResponse(c, berror.ErrPermissionDenied, nil)
		return
	}
	u.Role = m.Role

	// Validate the data.
	if err := u.Validate(); err != nil {
		log.ErrValidate(err)
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		log.ErrEncrypt(err)
		SendResponse(c, berror.ErrEncrypt, nil)
		return
	}

	// Save changed fields.
	if err := u.UpdateUser(); err != nil {
		log.ErrUpdateUser(err)
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}
	rsp := user.UpdateResponse{
		UserId:   u.ID,
		Username: u.Username,
	}
	SendResponse(c, nil, rsp)

}
