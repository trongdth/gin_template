package services

var (
	ErrInvalidEmail       = &Error{Code: -1000, Message: "invalid email"}
	ErrInvalidPassword    = &Error{Code: -1001, Message: "invalid password"}
	ErrEmailNotExists     = &Error{Code: -1002, Message: "email doesn't exist"}
	ErrEmailAlreadyExists = &Error{Code: -1003, Message: "email already exists"}
	ErrPasswordMismatch   = &Error{Code: -1004, Message: "password and confirm password must match"}
	ErrInvalidCredentials = &Error{Code: -1005, Message: "invalid credentials"}

	// general api errors
	ErrInvalidArgument     = &Error{Code: -9000, Message: "invalid argument"}
	ErrInternalServerError = &Error{Code: -9001, Message: "internal server error"}
	ErrInvalidLimit        = &Error{Code: -9002, Message: "invalid limit"}
	ErrInvalidPage         = &Error{Code: -9003, Message: "invalid page"}
	ErrSystemError         = &Error{Code: -9004, Message: "system error"}
	ErrPermissionDenied    = &Error{Code: -9005, Message: "permission denied"}
)

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}

func ErrorWithMessage(e *Error, message string) *Error {
	e.Message = message
	return e
}
