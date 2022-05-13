package common

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// SelfCheck returns self user information.
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

	SendResponse(c, nil, user)
}
