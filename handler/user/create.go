package user

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/berror"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	//log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, berror.ErrValidation, nil)
		return
	}

	// Validate if the user exists
	_, deleted, err := model.GetUser(r.Username)
	// if user data exists and deleted is false, send error
	//zap.L().Info("msg", zap.Bool("deleted", deleted), zap.Bool("err", err == nil))
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
	if err := u.Create(deleted); err != nil {
		SendResponse(c, berror.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
