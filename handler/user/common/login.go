package common

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/auth"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/pkg/token"
	"Jinshuzhai-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Login a user account
//
// @Summary Login  a user account
// @Description Login a user account with username and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.LoginRequest true "Login account"
// @Success 200 {object} user.SwaggerLoginResponse "{"code":0,"message":"OK","data":{"userId":7,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTA0NTkzODEsImlkIjo3LCJuYmYiOjE2NTA0NTkzODEsInJvbGUiOiJnZW5lcmFsIiwidXNlcm5hbWUiOiLkuIHno4oifQ.0kA4whaE9bZjXu4bN3Sw0DgrKwYzJ7kZenaGDOcdFRQ"}}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	zap.L().Info("Login function called.", zap.String("X-Request-Id", util.GetReqID(c)))
	// Binding the data with the user struct.
	var l user.LoginRequest
	//var u model.UserModel
	if err := c.Bind(&l); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, deleted, err := model.GetUser(l.Username)
	if deleted == true || err != nil {
		SendResponse(c, berror.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, l.Password); err != nil {
		SendResponse(c, berror.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(token.Context{ID: d.ID, Username: d.Username, Role: d.Role}, "")
	if err != nil {
		SendResponse(c, berror.ErrToken, nil)
		return
	}

	//SendResponse(c, nil, model.Token{Token: t})
	rsp := user.LoginResponse{
		UserId: d.ID,
		Token:  t,
	}

	SendResponse(c, nil, rsp)

}
