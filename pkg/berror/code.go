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
	ErrDeleteUser = &Berror{Code: 20106, Message: "Error occurred while deleting the user"}

	// ErrBookExists Book exists error
	ErrBookExists = &Berror{Code: 20201, Message: "The book was already exists."}
)
