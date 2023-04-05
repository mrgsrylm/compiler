package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/database"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/model"
)

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := model.Product{}

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	fmt.Println(product)

	product.UserID = uint(userData["id"].(float64))

	err = db.Create(&product).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	product := model.Product{}
	productID, _ := strconv.Atoi(ctx.Param("productID"))

	err := ctx.ShouldBindJSON(&product)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	product.UserID = uint(userData["id"].(float64))

	err = db.Model(&product).Where("id=?", productID).Updates(model.Product{Title: product.Title, Description: product.Description}).Error
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
