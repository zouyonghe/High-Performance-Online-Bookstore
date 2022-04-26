package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// RegisterSeller registers a new seller account.
//
// @Summary Register a new seller account.
// @Description Register a new seller account by username and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.RegisterRequest true "user information include username and password"
// @Success 200 {object} user.SwaggerRegisterResponse "{"code":0,"message":"OK","data":{"userId":12,"username":"汤桂英","role":"business"}}"
// @Router /user/register [post]
func RegisterSeller(c *gin.Context) {
	zap.L().Info("User create function called.", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	var r user.RegisterRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		zap.L().Error("Bind error.", zap.Error(err))
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Role:     "seller",
	}
	// Validate the data.
	if err := u.Validate(); err != nil {
		zap.L().Error("Error validating user data.", zap.Error(err))
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	// Validate if the user exists
	_, deleted, err := model.GetUser(r.Username)
	// if user data exists and deleted is false, send an error
	if deleted == false && err == nil {
		zap.L().Error("User already exists.", zap.String("username", r.Username))
		SendResponse(c, berror.ErrUserExists, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		zap.L().Error("Error encrypting user password.", zap.Error(err))
		SendResponse(c, berror.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.CreateUser(deleted); err != nil {
		zap.L().Error("Error creating user.", zap.Error(err))
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}

	rsp := user.RegisterResponse{
		UserId:   u.ID,
		Username: u.Username,
	}

	// Return the user id, username and role.
	SendResponse(c, nil, rsp)
}
