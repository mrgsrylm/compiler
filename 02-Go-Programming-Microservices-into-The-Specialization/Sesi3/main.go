package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gusrylmubarok/test/tree/main/02-Go-Programming-Microservices-into-The-Specialization/Sesi3/book"
	_ "github.com/lib/pq"
)

func main() {
	var PORT = ":4000"

	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:5432/book?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("database connection pool established +++++++++++++++++++++++++++++++++>")
	repo := book.NewRepository(db)
	service := book.NewService(repo)
	handler := book.NewHandler(service)

	server := gin.Default()
	server.GET("/book", handler.GetAll)
	server.GET("/book/:id", handler.GetById)
	server.POST("/book", handler.Create)
	server.PUT("/book/:id", handler.Update)
	server.DELETE("/book/:id", handler.DeleteById)

	server.Run(PORT)

}
