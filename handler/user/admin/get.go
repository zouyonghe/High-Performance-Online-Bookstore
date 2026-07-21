package admin

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Get gets a user account.
//
// @Summary Get a user information
// @Description Get a user account specified by user ID
// @Tags user/admin
// @Produce  json
// @Param id path uint64 true "the ID of the specified user to update"
// @Success 200 {object} user.SwaggerGetResponse "{"code":0,"message":"OK","data":{"UserID":5,"username":"傅秀英","role":"general"}}"
// @Router /user/admin/{id} [get]
// @Security ApiKeyAuth
func Get(c *gin.Context) {
	log.GetUserCalled(c)

	UserID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}
	// Get the user by the `id` from the database.
	u, err := model.GetUserByID(UserID)
	// if user is not found or deleted, send error
	if err != nil {
		SendError(c, err)
		return
	}

	// Return the public user information only,
	// never the password hash.
	SendResponse(c, nil, user.GetResponse{
		UserID:   u.ID,
		Username: u.Username,
		Role:     u.Role,
	})
}
