package log

import (
	"High-Performance-Online-Bookstore/util"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// User

func CheckPermissionCalled(c *gin.Context, sub, obj, act string) {
	zap.L().Debug("checkPermission", zap.String("sub", sub), zap.String("obj", obj), zap.String("act", act))

}

func RegisterCalled(c *gin.Context) {
	zap.L().Debug("create general user account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func LoginCalled(c *gin.Context) {
	zap.L().Debug("login function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func SelfDelCalled(c *gin.Context) {
	zap.L().Debug("delete self account function called", zap.String("X-Request-Id", util.GetReqID(c)))
}

func SelfUpdCalled(c *gin.Context) {
	zap.L().Debug("update self account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func SelfCheckCalled(c *gin.Context) {
	zap.L().Debug("check self account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DeleteUserCalled(c *gin.Context) {
	zap.L().Debug("delete function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func GetUserCalled(c *gin.Context) {
	zap.L().Debug("get user account information function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListUserCalled(c *gin.Context) {
	zap.L().Debug("list user function called", zap.String("X-Request-Id", util.GetReqID(c)))
}
func RegisterSellerCalled(c *gin.Context) {
	zap.L().Debug("create seller account function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func UpdateUserCalled(c *gin.Context) {
	zap.L().Debug("update user account information function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// Book

func AddBookCalled(c *gin.Context) {
	zap.L().Debug("add book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func UpdateBookCalled(c *gin.Context) {
	zap.L().Debug("update book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func GetBookCalled(c *gin.Context) {
	zap.L().Debug("get book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DelBookCalled(c *gin.Context) {
	zap.L().Debug("delete book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListBookCalled(c *gin.Context) {
	zap.L().Debug("list book function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// Cart

func ShowCartCalled(c *gin.Context) {
	zap.L().Debug("show cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func AddCartCalled(c *gin.Context) {
	zap.L().Debug("add book to cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DeleteCartCalled(c *gin.Context) {
	zap.L().Debug("delete book from cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ClearCartCalled(c *gin.Context) {
	zap.L().Debug("clear cart function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

// order

func CreateOrderCalled(c *gin.Context) {
	zap.L().Debug("create order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func DealOrderCalled(c *gin.Context) {
	zap.L().Debug("deal order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}

func ListOrderCalled(c *gin.Context) {
	zap.L().Debug("list order function called", zap.String("X-Request-Id", c.GetString("X-Request-Id")))
}
