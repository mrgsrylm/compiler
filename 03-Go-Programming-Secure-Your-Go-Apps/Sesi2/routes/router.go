package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/handler"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi2/middleware"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", handler.UserRegister)
		userRouter.POST("/login", handler.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middleware.Authentication())
		productRouter.POST("/", handler.CreateProduct)

		productRouter.PUT("/:productId", middleware.ProductAuthorization(), handler.UpdateProduct)
	}

	return r
}
