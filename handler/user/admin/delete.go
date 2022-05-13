package admin

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/service"
	"github.com/gin-gonic/gin"
)

// Delete deletes a user account.
//
// @Summary Delete a user
// @Description Delete a user by user ID
// @Tags user/admin
// @Accept  json
// @Produce  json
// @Param id path uint64 true "the ID of the specified user to delete"
// @Success 200 {object} user.SwaggerDeleteResponse "{"code":0,"message":"OK","data":{"UserID":5}}"
// @Router /user/admin/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	log.DeleteUserCalled(c)

	UserID, err := service.GetIDByParam(c)
	if err != nil {
		log.ErrParseToken(err)
		SendError(c, err)
		return
	}
	u, err := model.GetUserByID(UserID)
	if err != nil {
		log.ErrUserNotFound(err)
		SendError(c, err)
		return
	}
	username := u.Username
	if err = model.DeleteUser(UserID); err != nil {
		SendError(c, err)
		return
	}
	rsp := user.DeleteResponse{
		UserID:  UserID,
		Message: "User " + username + " deleted",
	}
	SendResponse(c, nil, rsp)
}
