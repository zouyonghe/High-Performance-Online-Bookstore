package handler

import (
	"High-Performance-Online-Bookstore/pkg/berror"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BaseResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := berror.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func SendDenyResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, gin.H{
		"message": "Forbidden",
	})
}

func SendError(c *gin.Context, err error) {
	SendResponse(c, err, nil)
}

func NoRoute(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"code":    http.StatusNotFound,
		"message": "The api does not exist.",
	})
}
