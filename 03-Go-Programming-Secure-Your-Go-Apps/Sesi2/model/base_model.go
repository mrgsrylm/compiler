package model

import "time"

type BaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at",omitempty`
	UpdatedAt time.Time `json:"updated_at",omitempty`
}
