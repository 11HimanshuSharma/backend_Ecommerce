package apperror
import "net/http"


type ErrorType int

const (
	// Bad Request - user sent invalid data (eg. nagative price)
	ErrorTypeBadRequest ErrorType = iota
	//not found - resource does not exist

	ErrorTypeNotFound
	// Conflict - business rule violation (eg, out of storkc)
	ErrorTypeConflict

	//Internal - something went wrong on our side
	ErrorTypeInternal

)

type AppError struct {
	Type ErrorType 
	Code string
	Message string 

}

func (e *AppError) Error() string {
	return e.Message
}


func (e *AppError) HTTPStatus() int {
	switch e.Type {
	case ErrorTypeBadRequest:
		return http.StatusBadRequest // 400
	case ErrorTypeNotFound:
		return http.StatusNotFound // 404
	case ErrorTypeConflict:
		return http.StatusConflict // 409
	case ErrorTypeInternal:
		return http.StatusInternalServerError // 500
	default:
		return http.StatusInternalServerError
	}

}

// constructor helpers
// These make creating errors clearn and consistent


func NewBadRequestError(code, message string ) *AppError {
	return &AppError{
		Type: ErrorTypeBadRequest,
		Code: code,
		Message: message,
	}
}

func NewNotFoundError(code, message string) *AppError {
	return &AppError{
		Type: ErrorTypeNotFound,
		Code: code,
		Message: message,
	}
}

func NewConflict(code, message string) *AppError {
	return &AppError {
		Type: ErrorTypeConflict,
		Code: code,
		Message: message,
	}
}


func NewInternalError(code, message string) *AppError {
	return &AppError{Type: ErrorTypeInternal, Code: code, Message: message}
}


func NewNotFound(code, message string) *AppError {
	return &AppError{
		Type: ErrorTypeNotFound,
		Code: code,
		Message: message,
	}
}


func NewInternal(code, message string) *AppError {
	return &AppError{
		Type: ErrorTypeInternal,
		Code: code,
		Message: message,
	}
}