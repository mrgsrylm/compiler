package module

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type productHandler struct {
	productUseCase ProductUseCase
}

func NewProductHandler(routers *gin.Engine, productUseCase ProductUseCase) {
	handler := &productHandler{productUseCase}

	router := routers.Group("api/v1/product")
	{
		router.POST("", handler.Insert)
		router.PUT("/:productId", handler.Update)
	}
}

func (h *productHandler) Insert(ctx *gin.Context) {
	var (
		product Product
		err     error
	)

	if err = ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	if err = h.productUseCase.Insert(ctx.Request.Context(), &product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, ResponseData{
		Status: "success",
		Data: &ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
		},
	})

}

func (h *productHandler) Update(ctx *gin.Context) {
	var (
		product    Product
		newProduct Product
		err        error
	)

	productID := ctx.Param("productId")

	if err = ctx.ShouldBindJSON(&product); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	updatedProduct := Product{
		ID:          productID,
		Title:       product.Title,
		Description: product.Description,
	}

	if newProduct, err = h.productUseCase.Update(ctx.Request.Context(), updatedProduct, productID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{
		Status: "success",
		Data: &ProductResponse{
			ID:          newProduct.ID,
			Title:       newProduct.Title,
			Description: newProduct.Description,
		},
	})
}
