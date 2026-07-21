package common

import (
	. "High-Performance-Online-Bookstore/handler"
	userpkg "High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// SelfCheck returns self user information.
//
// @Summary Get self user information
// @Description Get the information of the user specified by the token. The password hash is never returned.
// @Tags user/common
// @Produce json
// @Success 200 {object} user.SwaggerGetResponse
// @Router /user/common [get]
// @Security ApiKeyAuth
func SelfCheck(c *gin.Context) {
	log.SelfCheckCalled(c)

	UserID, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}

	user, err := model.GetUserByID(UserID)
	if err != nil {
		SendError(c, err)
		return
	}

	// Return the public user information only,
	// never the password hash.
	SendResponse(c, nil, userpkg.GetResponse{
		UserID:   user.ID,
		Username: user.Username,
		Role:     user.Role,
	})
}
