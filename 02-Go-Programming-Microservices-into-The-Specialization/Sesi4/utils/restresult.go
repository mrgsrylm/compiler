package utils

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Meta struct {
	Status  int    `json:"code"`
	Message string `json:"message"`
}

type Result struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func HandleSucces(c *gin.Context, data interface{}) {
	meta := Meta{
		Status:  200,
		Message: "success",
	}

	response := Result{
		Meta: meta,
		Data: data,
	}

	c.JSON(http.StatusOK, response)
}

func HandleError(c *gin.Context, status int, message string) {
	response := Meta{
		Status:  status,
		Message: message,
	}
	c.JSON(status, response)
}

func HandleValidator(c *gin.Context, data interface{}) {
	meta := Meta{
		Status:  422,
		Message: "error",
	}

	response := Result{
		Meta: meta,
		Data: data,
	}
	c.JSON(http.StatusUnprocessableEntity, response)
}

func ErrorValidators(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
