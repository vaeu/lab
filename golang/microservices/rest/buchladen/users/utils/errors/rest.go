package errors

import "net/http"

type RESTErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func NewBadRequest(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NewNotFound(message string) *RESTErr {
	return &RESTErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}
