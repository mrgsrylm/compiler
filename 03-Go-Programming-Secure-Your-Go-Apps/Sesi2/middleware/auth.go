package middleware

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/database"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/model"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/util"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := util.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthenticated",
				"message": err.Error(),
			})
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}

func ProductAuthorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := database.GetDB()
		productId, err := strconv.Atoi(c.Param("productId"))
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid Parameter",
			})
			return
		}
		userData := c.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := model.Product{}

		err = db.Select("user_id").First(&Product, uint(productId)).Error
		if err != nil {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "Data Doesn't Exist",
			})
			return
		}

		if Product.UserID != userID {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "You are not allowed to access this data",
			})
			return
		}
		c.Next()
	}
}
