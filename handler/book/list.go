package book

import (
	. "Jinshuzhai-Bookstore/handler"
	"Jinshuzhai-Bookstore/pkg/berror"
	"Jinshuzhai-Bookstore/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ListBook(c *gin.Context) {
	zap.L().Info("List book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
	var r ListRequest
	if err := c.ShouldBindJSON(&r); err != nil {
		zap.L().Error("Bind error", zap.Error(err))
		SendResponse(c, berror.ErrBind, nil)
		return
	}
	infos, count, err := service.ListBook(r.Title, r.PageNum, r.PageSize)
	if err != nil {
		zap.L().Error("List books error", zap.Error(err))
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		BookList:   infos,
	})
}
