package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi2/controllers"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/book", controllers.GetAllBooks)
	router.GET("/book/:id", controllers.GetBookById)
	router.POST("/book", controllers.CreateBook)
	router.PUT("/book/:id", controllers.UpdateBook)
	router.DELETE("/book/:id", controllers.DeleteBook)

	return router
}
