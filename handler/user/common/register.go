package common

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Register registers a new user account.
//
// @Summary Register a new user
// @Description Register a new user by username and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.RegisterRequest true "user information include username and password"
// @Success 200 {object} user.SwaggerRegisterResponse "{"code":0,"message":"OK","data":{"userId":12,"username":"汤桂英","role":"business"}}"
// @Router /user/register [post]
func Register(c *gin.Context) {
	zap.L().Info("User create function called.", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	var r user.RegisterRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	if r.Role == "" {
		r.Role = "general"
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Role:     r.Role,
	}
	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, berror.ErrValidation, nil)
		zap.L().Error("Error validating user data.", zap.Error(err))
		return
	}

	// Validate if the user exists
	_, deleted, err := model.GetUser(r.Username)
	// if user data exists and deleted is false, send error
	if deleted == false && err == nil {
		SendResponse(c, berror.ErrUserExists, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, berror.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.CreateUser(deleted); err != nil {
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}

	rsp := user.RegisterResponse{
		UserId:   u.ID,
		Username: u.Username,
		Role:     u.Role,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
