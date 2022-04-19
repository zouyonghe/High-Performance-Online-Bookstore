package user

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/model"
	"Jinshuzhai-Bookstore/pkg/auth"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		SendResponse(c, berror.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, deleted, err := model.GetUser(u.Username)
	if deleted == true || err != nil {
		SendResponse(c, berror.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, berror.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(token.Context{ID: d.ID, Username: d.Username, Role: d.Role}, "")
	if err != nil {
		SendResponse(c, berror.ErrToken, nil)
		return
	}

	SendResponse(c, nil, model.Token{Token: t})
}
