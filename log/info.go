package log

import (
	"High-Performance-Online-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// router

func StartListenHTTP(addr string) {
	zap.L().Info("start listening HTTP request", zap.String("addr", addr))
}

func StartListenHTTPS(addr string) {
	zap.L().Info("start listening HTTPS requests", zap.String("addr", addr))
}

func WaitForRouter() {
	zap.L().Info("Waiting for the router, retry in 1 second.")
}

func RouterDeployed() {
	zap.L().Info("The router has been deployed successfully.")
}

// User

func RegisterCalled(c *gin.Context) {
	zap.L().Info("create general user account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func LoginCalled(c *gin.Context) {
	zap.L().Info("login function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func SelfDelCalled(c *gin.Context) {
	zap.L().Info("delete self account function called", zap.String("X-Request-Id", util.GetReqID(c)))
}

func SelfUpdCalled(c *gin.Context) {
	zap.L().Info("update self account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func SelfCheckCalled(c *gin.Context) {
	zap.L().Info("check self account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DeleteUserCalled(c *gin.Context) {
	zap.L().Info("delete function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func GetUserCalled(c *gin.Context) {
	zap.L().Info("get user account information function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListUserCalled(c *gin.Context) {
	zap.L().Info("list user function called", zap.String("X-Request-Id", util.GetReqID(c)))
}
func RegisterSellerCalled(c *gin.Context) {
	zap.L().Info("create seller account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func UpdateUserCalled(c *gin.Context) {
	zap.L().Info("update user account information function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// Book

func AddBookCalled(c *gin.Context) {
	zap.L().Info("add book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func UpdateBookCalled(c *gin.Context) {
	zap.L().Info("update book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func GetBookCalled(c *gin.Context) {
	zap.L().Info("get book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DelBookCalled(c *gin.Context) {
	zap.L().Info("delete book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListBookCalled(c *gin.Context) {
	zap.L().Info("list book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// Cart

func ShowCartCalled(c *gin.Context) {
	zap.L().Info("show cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func AddCartCalled(c *gin.Context) {
	zap.L().Info("add book to cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DeleteCartCalled(c *gin.Context) {
	zap.L().Info("delete book from cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ClearCartCalled(c *gin.Context) {
	zap.L().Info("clear cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// order

func CreateOrderCalled(c *gin.Context) {
	zap.L().Info("create order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DealOrderCalled(c *gin.Context) {
	zap.L().Info("deal order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListOrderCalled(c *gin.Context) {
	zap.L().Info("list order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}
