package module

import (
	"fmt"
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
		router.DELETE("/:productId", handler.DeleteById)
		router.GET("", handler.FindAll)
		router.GET("/:productId", handler.FindById)
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

func (h *productHandler) DeleteById(ctx *gin.Context) {
	productID := ctx.Param("productId")

	if err := h.productUseCase.DeleteById(ctx.Request.Context(), productID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "your product has been successfully deleted",
	})
}

func (h *productHandler) FindAll(ctx *gin.Context) {
	var (
		products []Product
		err      error
	)

	if err = h.productUseCase.FindAll(ctx.Request.Context(), &products); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "faile",
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, ResponseData{
		Status: "success",
		Data:   products,
	})
}

func (h *productHandler) FindById(ctx *gin.Context) {
	var (
		product Product
		err     error
	)

	productID := ctx.Param("productId")

	fmt.Println(productID)

	if product, err = h.productUseCase.FindById(ctx.Request.Context(), productID); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, ResponseMessage{
			Status:  "fail",
			Message: err.Error(),
		})
		return
	}

	fmt.Println(product)

	ctx.JSON(http.StatusOK, ResponseData{
		Status: "success",
		Data: &ProductResponse{
			ID:          product.ID,
			Title:       product.Title,
			Description: product.Description,
		},
	})
}
