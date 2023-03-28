package book

import "time"

type Book struct {
	ID        int64     `gorm:"column:id;primaryKey"`
	Title     string    `gorm:"column:title;not null;unique;type:varchar(100)"`
	Author    string    `gorm:"column:author;not null;type:varchar(100)"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime:true"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdatedTime:true"`
}

type BookResponse struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type BookRequest struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
}

type GetBookDetailRequest struct {
	ID int `uri:"id" binding:"required"`
}

func (e *Book) TableName() string {
	return "books"
}
