package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi4/book"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi4/database"
)

func main() {
	var PORT = ":4000"

	database.StartDB()

	repo := book.NewRepository(database.GetDB())
	serv := book.NewService(repo)
	handler := book.NewHandler(serv)

	server := gin.Default()
	server.GET("/book", handler.GetAll)
	server.GET("/book/:id", handler.GetById)
	server.POST("/book", handler.Create)
	server.PUT("/book/:id", handler.Update)
	server.DELETE("/book/:id", handler.DeleteById)

	server.Run(PORT)
}
