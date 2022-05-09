package service

import (
	"High-Performance-Online-Bookstore/log"
	"High-Performance-Online-Bookstore/pkg/token"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetIDByParam(c *gin.Context) (uint64, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.ErrConv(err)
		return 0, err
	}
	return uint64(id), nil
}

func GetIDByToken(c *gin.Context) (uint64, error) {
	ctx, err := token.ParseRequest(c)
	if err != nil {
		return 0, err
	}
	return ctx.ID, nil
}

func GetRoleByToken(c *gin.Context) (string, error) {
	ctx, err := token.ParseRequest(c)
	if err != nil {
		return "guest", err
	}
	return ctx.Role, nil
}
