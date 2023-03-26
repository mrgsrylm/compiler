package utils

import "github.com/go-playground/validator/v10"

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type Result struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func RestResultNoData(message string, code int, status string) Meta {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	return meta
}

func RestResult(message string, code int, status string, data interface{}) Result {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	response := Result{
		Meta: meta,
		Data: data,
	}

	return response
}

func ErrorValidators(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
