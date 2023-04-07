package module

import (
	"context"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Product struct {
	ID          string `gorm:"column:id;primaryKey;type:VARCHAR(50)" json:"id"`
	Title       string `gorm:"column:title;not null;type:varchar(100)" valid:"required" json:"title"`
	Description string `gorm:"column:description;not null;type:varchar(100)" valid:"required" json:"description"`
}

func (p *Product) TableName() string {
	return "product"
}

func (p *Product) BeforeStore(db *gorm.DB) (err error) {
	if _, err := govalidator.ValidateStruct(p); err != nil {
		return err
	}
	return
}

type ProductRepository interface {
	Insert(context.Context, *Product) error
	Update(context.Context, Product, string) (Product, error)
	// DeleteById(context.Context, string) error
	// FindAll(context.Context, *[]Product) error
	// FindById(context.Context, *Product, string)
}

type ProductUseCase interface {
	Insert(context.Context, *Product) error
	Update(context.Context, Product, string) (Product, error)
	// DeleteById(context.Context, string) error
	// FindAll(context.Context, *[]Product) error
	// FindById(context.Context, *Product, string)
}

type ProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type ProductResponse struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
