package admin

import (
	. "High-Performance-Online-Bookstore/handler"
	"High-Performance-Online-Bookstore/handler/user"
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/model"
	"High-Performance-Online-Bookstore/pkg/berror"
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
// @Success 200 {object} user.SwaggerDeleteResponse "{"code":0,"message":"OK","data":{"userId":5}}"
// @Router /user/admin/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	log.DeleteUserCalled(c)

	userId, err := service.GetIDByToken(c)
	if err != nil {
		log.ErrParseToken(err)
		SendResponse(c, berror.InternalServerError, nil)
	}
	um, err := model.GetUserByID(userId)
	if err != nil {
		log.ErrUserNotFound(err)
		c.JSON(400, gin.H{
			"message": "User not found",
		})
		return
	}
	username := um.Username
	if err := model.DeleteUser(userId); err != nil {
		SendResponse(c, err, nil)
		return
	}
	rsp := user.DeleteResponse{
		UserId:  userId,
		Message: "User " + username + " deleted",
	}
	SendResponse(c, nil, rsp)
}
