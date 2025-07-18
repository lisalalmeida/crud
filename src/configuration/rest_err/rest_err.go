package rest_err

import (
	"net/http"
)


type RestErr struct{
	Message string `json:"message"`
	Err string `json:"error"`
	Code int `json:"code"`
	Causes []Causes `json:"causes"`
}

type Causes struct{
	Field string `json:"field"`
	Message string `json:"message"`
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: err,
		Code: code,
		Causes: causes,
	}
}

func NewBandRequestError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
	}
}

func NewBandRequestValidationError(message string, causes []Causes) *RestErr{
	return &RestErr{
		Message: message,
		Err: "bad_request",
		Code: http.StatusBadRequest,
		Causes: causes,
	}
}

func NewInternalServerError(message string) *RestErr{
	return &RestErr{
		Message: message,
		Err: "internal_server_error",
		Code: http.StatusInternalServerError,
	}
}

