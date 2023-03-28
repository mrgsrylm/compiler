package book

import (
	"fmt"
	"gorm.io/gorm"
)

type Repository interface {
	FindAll() (*[]Book, error)
	FindById(id int) (*Book, error)
	Insert(book *Book) (*Book, error)
	Update(id int, book *Book) (*Book, error)
	DeleteById(id int) error
}

type repo struct {
	DB *gorm.DB
}

func NewRepository(DB *gorm.DB) *repo {
	return &repo{DB}
}

func (r *repo) FindAll() (*[]Book, error) {
	var books []Book

	err := r.DB.Find(&books).Error
	if err != nil {
		fmt.Printf("[repo.FindAll] error execute query %v \n", err)
		return nil, err
	}

	return &books, nil
}

func (r *repo) FindById(id int) (*Book, error) {
	var book = Book{}

	err := r.DB.Table("books").Where("id = ?", id).First(&book).Error
	if err != nil {
		fmt.Printf("[repo.FindById] error execute query %v \n", err)
		return nil, fmt.Errorf("id is not exsis")
	}

	return &book, nil
}

func (r *repo) Insert(book *Book) (*Book, error) {
	err := r.DB.Create(&book).Error
	if err != nil {
		fmt.Printf("[repo.Insert] error execute query %v \n", err)
		return book, err
	}

	return book, nil
}

func (r *repo) Update(id int, book *Book) (*Book, error) {
	var newBook = Book{} 
	err := r.DB.Table("books").Where("id = ?", id).First(&newBook).Updates(&book).Error
	if err != nil {
		fmt.Printf("[repo.Update] error execute query %v \n", err)
		return book, err
	}

	return &newBook, nil
}

func (r *repo) DeleteById(id int) error {
	var book = Book{}

	err := r.DB.Table("books").Where("id = ?", id).First(&book).Delete(&book).Error
	if err != nil {
		fmt.Printf("[repo.Delete] error execute query %v \n", err)
		return fmt.Errorf("id is not exsis")
	}

	return nil
}
