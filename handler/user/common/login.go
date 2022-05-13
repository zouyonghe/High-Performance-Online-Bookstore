package common

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/auth"
	"High-Performance-Online-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
)

// Login a user account
//
// @Summary Login  a user account
// @Description Login a user account with username and password
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.LoginRequest true "Login account"
// @Success 200 {object} user.SwaggerLoginResponse "{"code":0,"message":"OK","data":{"UserID":7,"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE2NTA0NTkzODEsImlkIjo3LCJuYmYiOjE2NTA0NTkzODEsInJvbGUiOiJnZW5lcmFsIiwidXNlcm5hbWUiOiLkuIHno4oifQ.0kA4whaE9bZjXu4bN3Sw0DgrKwYzJ7kZenaGDOcdFRQ"}}"
// @Router /user/login [post]
func Login(c *gin.Context) {
	log.LoginCalled(c) // Binding the data with the user struct.

	// bind request body
	var r user.LoginRequest
	//var u model.UserModel
	if err := c.ShouldBindJSON(&r); err != nil {
		log.ErrBind(err)
		SendError(c, err)
		return
	}

	// Get the user information by the login username.
	d, deleted, err := model.GetUser(r.Username)
	if deleted == true || err != nil {
		SendError(c, err)
		return
	}

	// Compare the login password with the user password.
	if err = auth.Compare(d.Password, r.Password); err != nil {
		SendError(c, err)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(token.Context{ID: d.ID, Username: d.Username, Role: d.Role}, "")
	if err != nil {
		SendError(c, err)
		return
	}

	//SendResponse(c, nil, model.Token{Token: t})
	rsp := user.LoginResponse{
		UserID: d.ID,
		Token:  t,
	}

	SendResponse(c, nil, rsp)

}
