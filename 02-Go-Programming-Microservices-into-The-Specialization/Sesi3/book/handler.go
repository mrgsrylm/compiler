package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi3/utils"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service}
}

func (h *handler) GetAll(ctx *gin.Context) {
	books, err := h.service.GetAll()
	if err != nil {
		res := utils.RestResultNoData("failed to get books", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.RestResult("get all books", http.StatusOK, "success", ToBooks(books))
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) GetById(ctx *gin.Context) {
	var req GetBookDetailRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		res := utils.RestResultNoData("failed to get book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	book, err := h.service.GetById(int64(req.ID))
	if err != nil {
		res := utils.RestResultNoData("failed to get book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.RestResult("get all book", http.StatusOK, "success", ToBook(book))
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) Create(ctx *gin.Context) {
	var req BookRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errors := utils.ErrorValidators(err)
		msg := gin.H{"errors": errors}

		res := utils.RestResult("validation failed", http.StatusUnprocessableEntity, "error", msg)
		ctx.JSON(http.StatusUnprocessableEntity, res)
		return
	}

	newBook, err := h.service.Create(ToEntity(req))
	if err != nil {
		res := utils.RestResultNoData("failed to create book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	res := utils.RestResult("book", http.StatusOK, "success", ToBook(newBook))
	ctx.JSON(http.StatusOK, res)
}

func (h *handler) Update(ctx *gin.Context) {
	var reqId GetBookDetailRequest
	err := ctx.ShouldBindUri(&reqId)
	if err != nil {
		response := utils.RestResultNoData("failed to update book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	var reqData BookRequest
	err = ctx.ShouldBindJSON(&reqData)
	if err != nil {
		errors := utils.ErrorValidators(err)
		errorMessage := gin.H{"errors": errors}

		response := utils.RestResult("validation failed", http.StatusUnprocessableEntity, "error", errorMessage)
		ctx.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	newBook, err := h.service.Update(int64(reqId.ID), ToEntity(reqData))
	if err != nil {
		response := utils.RestResultNoData("failed to update book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.RestResult("success to update book", http.StatusOK, "success", ToBook(newBook))
	ctx.JSON(http.StatusOK, response)
}

func (h *handler) DeleteById(ctx *gin.Context) {
	var reqId GetBookDetailRequest
	err := ctx.ShouldBindUri(&reqId)
	if err != nil {
		response := utils.RestResultNoData("failed to delete book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = h.service.DeleteById(int64(reqId.ID))
	if err != nil {
		response := utils.RestResultNoData("failed to delete book", http.StatusBadRequest, "error")
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	response := utils.RestResult("success to delete book", http.StatusOK, "success", nil)
	ctx.JSON(http.StatusOK, response)
}
