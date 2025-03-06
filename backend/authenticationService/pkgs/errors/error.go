package cError

import "net/http"

// List of all error
const (
	InternalServerError     = "Internal Server Error"
	InvalidRequestError     = "Invalid Request"
	ServiceUnavailableError = "Service Unavailable"
	NotFoundError           = "Not Found"
	BadRequestError        = "Bad Request"
	UnauthorizedError      = "Unauthorized"
	ConflictError          = "Conflict"
)

const (
	ServiceUnavailableCode = http.StatusServiceUnavailable
	InternalServerCode     = http.StatusInternalServerError
	InvalidRequestCode     = http.StatusBadRequest
	NotFoundCode           = http.StatusNotFound
	BadRequestCode        = http.StatusBadRequest
	UnauthorizedCode      = http.StatusUnauthorized
	ConflictCode          = http.StatusConflict  // 409
)

type IError interface {
	Error() string
	GetHttpCode() int
}

type CustomError struct {
	Code    int    `json:"-"`
	Status  string `json:"status"`
	Message string `json:"message"`

	// Errors should be used to return validation errors
	ValidationError []ValidationErrorField `json:"validation_error,omitempty"`
}

type ValidationErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (c *CustomError) Error() string {
	return c.Message
}

func (c *CustomError) GetHttpCode() int {
	return c.Code
}

func getErrorCode(errMessage string) int {
	var val int

	switch errMessage {
	case ServiceUnavailableError:
		val = ServiceUnavailableCode
	case InvalidRequestError:
		val = InvalidRequestCode
	case NotFoundError:
		val = NotFoundCode
	case BadRequestError:
		val = BadRequestCode
	case UnauthorizedError:
		val = UnauthorizedCode
	case ConflictError:
		val = ConflictCode
	default:
		val = InternalServerCode
	}
	return val
}

func GetError(errMessage string, errDetail error) *CustomError {
	return &CustomError{
		Code:    getErrorCode(errMessage),
		Status:  errMessage,
		Message: errDetail.Error(),
	}
}

func GetErrorValidation(errMessage string, errDetail error, validationError []ValidationErrorField) *CustomError {
	return &CustomError{
		Code:            getErrorCode(errMessage),
		Status:          errMessage,
		Message:         errDetail.Error(),
		ValidationError: validationError,
	}
}
