package berror

// ---- success ----
var (
	// OK Success with no errors
	OK = &Berror{Code: 0, Message: "OK"}
)

// ---- fail ----

//common
var (
	// InternalServerError Internal server error
	InternalServerError = &Berror{Code: 10001, Message: "Internal server error."}
	// ErrBindRequest binding the request body to the struct failed
	ErrBindRequest = &Berror{Code: 10002, Message: "Error occurred while binding the request body."}
	// ErrPermissionDenied Permission denied
	ErrPermissionDenied = &Berror{Code: 10003, Message: "Permission denied"}
	// ErrSignToken sign token failed
	ErrSignToken = &Berror{Code: 10004, Message: "Error occurred while signing the token."}
	// ErrParseToken parse token failed
	ErrParseToken = &Berror{Code: 10005, Message: "Parse token failed"}
	// ErrTokenInvalid token was invalid error
	ErrTokenInvalid = &Berror{Code: 10006, Message: "The token was invalid."}
	// ErrValidation validation failed
	ErrValidation = &Berror{Code: 10007, Message: "Validation failed."}
	// ErrDatabase database operation failed
	ErrDatabase = &Berror{Code: 10008, Message: "Database error."}
	// ErrEncrypt encrypting the user password failed
	ErrEncrypt = &Berror{Code: 10109, Message: "Error occurred while encrypting the user password."}
)

//user
var (
	// ErrCreateUser create user failed
	ErrCreateUser = &Berror{Code: 20101, Message: "Error occurred while creating the user."}
	// ErrUserNotFound user was not found
	ErrUserNotFound = &Berror{Code: 20102, Message: "The user was not found."}
	// ErrUserExists user already exists
	ErrUserExists = &Berror{Code: 20103, Message: "The username is already exists."}
	// ErrPasswordIncorrect user password was incorrect
	ErrPasswordIncorrect = &Berror{Code: 20105, Message: "The password was incorrect."}
	// ErrDeleteUser deletes the user failed
	ErrDeleteUser = &Berror{Code: 20106, Message: "Error occurred while deleting the user."}
)

//book
var (
	// ErrCreateBook creates a book failed
	ErrCreateBook = &Berror{Code: 20200, Message: "Error occurred while creating a book"}
	// ErrBookExists book exists error
	ErrBookExists = &Berror{Code: 20201, Message: "The book was already exists."}
	// ErrGetBook book was not found
	ErrGetBook = &Berror{Code: 20202, Message: "The book was not found."}
	// ErrGetBookList get book list failed
	ErrGetBookList = &Berror{Code: 20203, Message: "Error occurred while getting book list."}
	// ErrBookNotExist book was not found
	ErrBookNotExist = &Berror{Code: 20204, Message: "The book is not exist."}
	// ErrBookNotEnough book not enough
	ErrBookNotEnough = &Berror{Code: 20205, Message: "The book is not enough."}
	// ErrBookNotSell book not sell
	ErrBookNotSell = &Berror{Code: 20206, Message: "The book is not sell."}
	// ErrBookInCartNotEnough book in cart not enough
	ErrBookInCartNotEnough = &Berror{Code: 20207, Message: "The book in cart is not enough."}
)

//cart
var (
	// ErrAddCart add cart failed
	ErrAddCart = &Berror{Code: 20300, Message: "Error occurred while adding the book to cart."}
	// ErrGetCart get cart failed
	ErrGetCart = &Berror{Code: 20301, Message: "Error occurred while getting the cart."}
	// ErrDeleteCart delete cart failed
	ErrDeleteCart = &Berror{Code: 20302, Message: "Error occurred while deleting the cart."}
	// ErrUpdateCart update cart failed
	ErrUpdateCart = &Berror{Code: 20303, Message: "Error occurred while updating the cart."}
	// ErrClearCart clear cart failed
	ErrClearCart = &Berror{Code: 20205, Message: "Error occurred while clearing the cart."}
	// ErrNothingInCart nothing in cart
	ErrNothingInCart = &Berror{Code: 20206, Message: "Nothing in cart."}
	// ErrDeleteBookFromCart delete book from cart failed
	ErrDeleteBookFromCart = &Berror{Code: 20207, Message: "Error occurred while deleting the book from cart."}
)

//order
var (
	// ErrCreateOrder create order failed
	ErrCreateOrder = &Berror{Code: 20400, Message: "Error occurred while creating the order."}
	// ErrGetOrder get order failed
	ErrGetOrder = &Berror{Code: 20401, Message: "Error occurred while getting the order."}
	// ErrGetOrderList get order list failed
	ErrGetOrderList = &Berror{Code: 20402, Message: "Error occurred while getting the order list."}
	// ErrDeleteOrder delete order failed
	ErrDeleteOrder = &Berror{Code: 20403, Message: "Error occurred while deleting the order."}
	// ErrUpdateOrder update order failed
	ErrUpdateOrder = &Berror{Code: 20404, Message: "Error occurred while updating the order."}
	// ErrOrderNotExist order was not found
	ErrOrderNotExist = &Berror{Code: 20405, Message: "The order is not exist."}
	// ErrApproveOrder approve order failed
	ErrApproveOrder = &Berror{Code: 20210, Message: "Error occurred while approving the order."}
	// ErrDealOrder deal order failed
	ErrDealOrder = &Berror{Code: 20211, Message: "Error occurred while dealing the order."}
	// ErrAddBookToOrder add book to order failed
	ErrAddBookToOrder = &Berror{Code: 20212, Message: "Error occurred while adding the book to order."}
	// ErrSetOrderPrice set order price failed
	ErrSetOrderPrice = &Berror{Code: 20213, Message: "Error occurred while setting the order price."}
)
