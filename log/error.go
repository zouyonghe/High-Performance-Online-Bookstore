package log

import "go.uber.org/zap"

func ErrListenHTTP(err error) {
	zap.L().Error("error listening HTTP request", zap.Error(err))
}

func ErrListenHTTPS(err error) {
	zap.L().Error("error listening HTTPS request", zap.Error(err))
}

func ErrNoResponse(err error) {
	zap.L().Error("The router has no response, or it might took too long to start up.", zap.Error(err))
}

func ErrConv(err error) {
	zap.L().Error("error converting type", zap.Error(err))
}

func ErrParseToken(err error) {
	zap.L().Error("error parsing token", zap.Error(err))
}

func ErrBind(err error) {
	zap.L().Error("error binding", zap.Error(err))
}

func ErrValidate(err error) {
	zap.L().Error("error validating", zap.Error(err))
}

func ErrEncrypt(err error) {
	zap.L().Error("error encrypting", zap.Error(err))
}

// User error

func ErrCreateUser(err error) {
	zap.L().Error("error creating user", zap.Error(err))
}

func ErrUserExists() {
	zap.L().Error("user already exists error")
}

func ErrDeleteUser(err error) {
	zap.L().Error("error deleting user", zap.Error(err))
}
func ErrGetUser(err error) {
	zap.L().Error("error getting user", zap.Error(err))
}

func ErrListUsers(err error) {
	zap.L().Error("error listing users error", zap.Error(err))
}

func ErrUserNotFound(err error) {
	zap.L().Error("user not found error", zap.Error(err))
}

func ErrUpdateUser(err error) {
	zap.L().Error("error updating user", zap.Error(err))
}

// Book

func ErrBookExists() {
	zap.L().Error("book already exists error")
}

func ErrCreateBook(err error) {
	zap.L().Error("error creating book", zap.Error(err))
}

func ErrGetBook(err error) {
	zap.L().Error("error getting book", zap.Error(err))
}

func ErrDelBook(err error) {
	zap.L().Error("error deleting book", zap.Error(err))
}

func ErrListBooks(err error) {
	zap.L().Error("error listing books", zap.Error(err))
}

func ErrUpdateBook(err error) {
	zap.L().Error("error updating book", zap.Error(err))
}

func ErrCreateCart(err error) {
	zap.L().Error("error creating cart", zap.Error(err))
}

func ErrAddCart(err error) {
	zap.L().Error("error adding book to cart", zap.Error(err))
}

func ErrCreateCartBook(err error) {
	zap.L().Error("error creating cartBook", zap.Error(err))
}

func ErrGetCart(err error) {
	zap.L().Error("error getting cart", zap.Error(err))
}

func ErrDeleteCart(err error) {
	zap.L().Error("error deleting cart", zap.Error(err))
}

func ErrClearCart(err error) {
	zap.L().Error("error clearing cart", zap.Error(err))
}

func ErrGetCartBookList(err error) {
	zap.L().Error("error getting cartBookList", zap.Error(err))
}
func ErrCreateOrder(err error) {
	zap.L().Error("error creating order", zap.Error(err))
}

func ErrGetOrder(err error) {
	zap.L().Error("error getting order", zap.Error(err))
}

func ErrGetOrderBook(err error) {
	zap.L().Error("error getting orderBook", zap.Error(err))
}

func ErrAddOrder(err error) {
	zap.L().Error("error adding book to order", zap.Error(err))
}

func ErrUpdateOrderPrice(err error) {
	zap.L().Error("error updating order price", zap.Error(err))
}

func ErrApproveOrder(err error) {
	zap.L().Error("error approving order", zap.Error(err))
}

func ErrDeleteOrder(err error) {
	zap.L().Error("error deleting order", zap.Error(err))
}

func ErrDealOrder(err error) {
	zap.L().Error("error dealing order", zap.Error(err))
}
