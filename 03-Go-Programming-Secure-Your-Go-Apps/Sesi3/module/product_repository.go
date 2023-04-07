package module

import (
	"context"
	"fmt"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type productRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *productRepository {
	return &productRepository{db}
}

func (p *productRepository) Insert(ctx context.Context, product *Product) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ID, _ := gonanoid.New(16)

	product.ID = fmt.Sprintf("product-%s", ID)

	if err = p.db.WithContext(ctx).Create(&product).Error; err != nil {
		return err
	}

	return
}

func (p *productRepository) Update(ctx context.Context, newProduct Product, id string) (product Product, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	product = Product{}

	if err = p.db.WithContext(ctx).Model(&product).Where("id = ?", &id).Updates(newProduct).Error; err != nil {
		return product, err
	}

	if err = p.db.WithContext(ctx).First(&product, "id = ?", product.ID).Error; err != nil {
		return product, err
	}

	return product, nil
}
