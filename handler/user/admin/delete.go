package admin

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/handler/user"
	"Jinshuzhai-Bookstore/model"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// Delete deletes a user account.
//
// @Summary Delete a user
// @Description Delete a user by user ID
// @Tags user/admin
// @Accept  json
// @Produce  json
// @Param id path uint64 true "the ID of the specified user to delete"
// @Success 200 {object} user.DeleteResponse "{"code":0,"message":"OK","data":{"userId":5}}"
// @Router /user/admin/{id} [delete]
// @Security ApiKeyAuth
func Delete(c *gin.Context) {
	zap.L().Info("delete function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	userId, _ := strconv.Atoi(c.Param("id"))
	if _, err := model.GetUserByID(uint64(userId)); err != nil {
		SendResponse(c, err, nil)
		return
	}
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, err, nil)
		return
	}
	rsp := user.DeleteResponse{
		UserId: uint64(userId),
	}
	SendResponse(c, nil, rsp)
}
