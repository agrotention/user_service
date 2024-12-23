package errors

var (
	InternalError = NewServiceError(500, "internal error")
	UserNotFound  = NewServiceError(404, "user not found")
)

type ServiceError struct {
	httpCode int
	message  string
}

func (e *ServiceError) Error() string {
	return e.message
}

func NewServiceError(httpCode int, message string) *ServiceError {
	return &ServiceError{httpCode, message}
}
