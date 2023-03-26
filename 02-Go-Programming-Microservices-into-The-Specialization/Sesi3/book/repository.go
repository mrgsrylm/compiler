package book

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"time"
)

type Repository interface {
	FindAll() ([]Book, error)
	FindById(id int64) (Book, error)
	Insert(book Book) (int64, error)
	Update(book Book) (int64, error)
	DeleteById(id int64) (int64, error)
}

type repo struct {
	DB *sql.DB
}

func NewRepository(DB *sql.DB) *repo {
	return &repo{DB}
}

func (r *repo) FindAll() ([]Book, error) {
	query := `
		SELECT * FROM book
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		log.Println("cannot query")
		return nil, (errors.New("cannot run query"))
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		var book Book

		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.DescBook)
		if err != nil {
			log.Println("cannot fetch data")
			return nil, errors.New("cannot fetch all book")
		}

		books = append(books, book)
	}

	return books, err
}

func (r *repo) FindById(id int64) (Book, error) {
	var book Book

	query := `
		SELECT id, title, author, desc_book
		FROM book
		WHERE id=$1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&book.ID,
		&book.Title,
		&book.Author,
		&book.DescBook,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return book, errors.New("record not found")
		default:
			return book, err
		}
	}

	return book, nil
}

func (r *repo) Insert(book Book) (int64, error) {
	query := `
		INSERT INTO book (title, author, desc_book)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	args := []interface{}{book.Title, book.Author, book.DescBook}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return id, errors.New("record not found")
		default:
			return id, err
		}
	}

	return id, nil
}

func (r *repo) Update(book Book) (int64, error) {
	query := `
		UPDATE book
		SET title = $1, author = $2, desc_book = $3
		WHERE id = $4
		RETURNING id
	`

	args := []interface{}{book.Title, book.Author, book.DescBook, book.ID}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var id int64

	err := r.DB.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return id, errors.New("update failed")
		default:
			return id, err
		}
	}

	return id, nil
}

func (r *repo) DeleteById(id int64) (int64, error) {
	if id < 1 {
		return 0, errors.New("record not found")
	}

	query := `
		DELETE FROM book
		WHERE id = $1
	`

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	result, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return 0, errors.New("cannot run query")
	}

	rowAffected, err := result.RowsAffected()
	if err != nil {
		return 0, errors.New("cannot run query")
	}
	if rowAffected == 0 {
		return 0, errors.New("record not found")
	}

	return rowAffected, nil
}
