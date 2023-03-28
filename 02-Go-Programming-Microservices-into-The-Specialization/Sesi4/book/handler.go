package book

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi4/utils"
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
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.HandleSucces(ctx, ToBooks(*books))
}

func (h *handler) GetById(ctx *gin.Context) {
	var req GetBookDetailRequest

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	book, err := h.service.GetById(req.ID)
	if err != nil {
		utils.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSucces(ctx, ToBook(*book))
}

func (h *handler) Create(ctx *gin.Context) {
	var req BookRequest

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		errors := utils.ErrorValidators(err)
		msg := gin.H{"errors": errors}
		utils.HandleValidator(ctx, msg)
		return
	}


	book, err := h.service.Create(req)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSucces(ctx, ToBook(*book))
}

func (h *handler) Update(ctx *gin.Context) {
	var reqParam GetBookDetailRequest

	err := ctx.ShouldBindUri(&reqParam)
	if err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	var reqBody BookRequest
	err = ctx.ShouldBindJSON(&reqBody)
	if err != nil {
		errors := utils.ErrorValidators(err)
		msg := gin.H{"errors": errors}
		utils.HandleValidator(ctx, msg)
		return
	}

	newBook, err := h.service.Update(reqParam.ID, reqBody)
	if err != nil {
		utils.HandleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.HandleSucces(ctx, ToBook(*newBook))
}

func (h *handler) DeleteById(ctx *gin.Context) {
	var reqParam GetBookDetailRequest

	err := ctx.ShouldBindUri(&reqParam)
	if err != nil {
		utils.HandleError(ctx, http.StatusBadRequest, "id has be number")
		return
	}

	err = h.service.DeleteById(reqParam.ID)
	if err != nil {
		utils.HandleError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.HandleSucces(ctx, "data has been deleted")
}
