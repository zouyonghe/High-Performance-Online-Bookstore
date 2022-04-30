package common

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
	"High-Performance-Online-Bookstore/service"
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
// @Success 200 {object} user.SwaggerSelfUpdResponse "{"code":0,"message":"OK","data":{"UserID":6,"username":"夏秀兰"}}"
// @Router /user/common [put]
// @Security ApiKeyAuth
func SelfUpd(c *gin.Context) {
	log.SelfUpdCalled(c)
	UserID, err := service.GetIDByToken(c)
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
	u, err := model.GetUserByID(UserID)
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
		UserID:   UserID,
		Username: u.Username,
	}
	//SendResponse(c, nil, nil)
	SendResponse(c, nil, rsp)
}
