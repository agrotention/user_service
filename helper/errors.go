package helper

import "log"

type ServiceError struct {
	httpCode int
	message  string
}

func (e *ServiceError) Error() string {
	return e.message
}

func NewServiceError(httpCode int, message string, source ...any) *ServiceError {
	log.Println(message, source)
	return &ServiceError{httpCode, message}
}
