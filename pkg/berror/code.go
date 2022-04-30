package berror

var (
	// OK Success with no errors
	OK = &Berror{Code: 0, Message: "OK"}

	// InternalServerError Internal server error
	InternalServerError = &Berror{Code: 10001, Message: "Internal server error"}
	// ErrBind binding the request body to the struct failed
	ErrBind = &Berror{Code: 10002, Message: "Error occurred while binding the request body to the struct."}
	// ErrPermissionDenied Permission denied
	ErrPermissionDenied = &Berror{Code: 10003, Message: "Permission denied"}

	// ErrValidation validation failed
	ErrValidation = &Berror{Code: 20001, Message: "Validation failed."}
	// ErrDatabase database operation failed
	ErrDatabase = &Berror{Code: 20002, Message: "Database error."}
	// ErrToken signing the JSON web token failed
	ErrToken = &Berror{Code: 20003, Message: "Error occurred while signing the JSON web token."}
	// ErrCheckRole check the user role failed
	ErrCheckRole = &Berror{Code: 20004, Message: "Error occurred while checking the role."}

	// ErrEncrypt encrypting the user password failed
	ErrEncrypt = &Berror{Code: 20101, Message: "Error occurred while encrypting the user password."}
	// ErrUserNotFound user was not found
	ErrUserNotFound = &Berror{Code: 20102, Message: "The user was not found."}
	// ErrUserExists user already exists
	ErrUserExists = &Berror{Code: 20103, Message: "The username is already exists."}
	// ErrTokenInvalid token was invalid error
	ErrTokenInvalid = &Berror{Code: 20104, Message: "The token was invalid."}
	// ErrPasswordIncorrect user password was incorrect
	ErrPasswordIncorrect = &Berror{Code: 20105, Message: "The password was incorrect."}
	// ErrDeleteUser deletes the user failed
	ErrDeleteUser = &Berror{Code: 20106, Message: "Error occurred while deleting the user."}

	// ErrBookExists book exists error
	ErrBookExists = &Berror{Code: 20201, Message: "The book was already exists."}
	// ErrGetBook book was not found
	ErrGetBook = &Berror{Code: 20202, Message: "The book was not found."}
	// ErrGetBookList get book list failed
	ErrGetBookList = &Berror{Code: 20203, Message: "Error occurred while getting book list."}
	// ErrBookNotExist book was not found
	ErrBookNotExist = &Berror{Code: 20204, Message: "The book is not exist."}
	// ErrClearCart clear cart failed
	ErrClearCart = &Berror{Code: 20205, Message: "Error occurred while clearing the cart."}
	// ErrNothingInCart nothing in cart
	ErrNothingInCart = &Berror{Code: 20206, Message: "Nothing in cart."}
)
