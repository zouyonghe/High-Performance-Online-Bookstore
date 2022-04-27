package common

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/log"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// SelfUpd updates the current user information
//
// @Summary Update the current user information
// @Description Update the current user information by username and password
// @Tags user/common
// @Accept  json
// @Produce  json
// @Param user body user.SelfUpdRequest true "Create a new user"
// @Success 200 {object} user.SwaggerSelfUpdResponse "{"code":0,"message":"OK","data":{"userId":6,"username":"夏秀兰"}}"
// @Router /user/common [put]
// @Security ApiKeyAuth
func SelfUpd(c *gin.Context) {
	log.SelfUpdCalled(c)
	userId, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
	}

	var r user.SelfUpdRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	u, err := model.GetUserByID(userId)
	if err != nil {
		log.ErrGetUser(err)
		SendResponse(c, berror.ErrDatabase, nil)
	}
	if r.Username != "" {
		u.Username = r.Username
	}
	if r.Password != "" {
		u.Password = r.Password
	}

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
	if err := u.UpdateUser(); err != nil {
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}
	rsp := user.SelfUpdResponse{
		UserId:   userId,
		Username: u.Username,
	}
	//SendResponse(c, nil, nil)
	SendResponse(c, nil, rsp)
}
