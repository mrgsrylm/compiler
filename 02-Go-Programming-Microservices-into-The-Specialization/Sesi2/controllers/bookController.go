package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func GetAllBooks(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": BookDatas,
	})
}

func GetBookById(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if id == book.ID {
			condition = true
			bookData = BookDatas[i]
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookData,
	})
}

func CreateBook(ctx *gin.Context) {
	var newBook Book
	if err := ctx.ShouldBindJSON(&newBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	newBook.ID = fmt.Sprintf("%d", len(BookDatas)*1)
	BookDatas = append(BookDatas, newBook)

	ctx.JSON(http.StatusCreated, gin.H{
		"data": newBook,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	for i, book := range BookDatas {
		if id == book.ID {
			condition = true
			BookDatas[i] = updatedBook
			BookDatas[i].ID = id
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully updated", id),
	})
}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookIdx int

	for i, book := range BookDatas {
		if id == book.ID {
			condition = true
			bookIdx = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Data Not Found",
			"error_message": fmt.Sprintf("book with id %v not found", id),
		})
		return
	}

	copy(BookDatas[bookIdx:], BookDatas[bookIdx+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("book with id %v has been successfully deleted", id),
	})
}
