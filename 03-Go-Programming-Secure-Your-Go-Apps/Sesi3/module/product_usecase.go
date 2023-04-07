package module

import (
	"context"
)

type productUseCase struct {
	productRepository ProductRepository
}

func NewProductUseCase(productRepository ProductRepository) *productUseCase {
	return &productUseCase{productRepository}
}

func (u *productUseCase) Insert(ctx context.Context, product *Product) (err error) {
	if err = u.productRepository.Insert(ctx, product); err != nil {
		return err
	}
	return
}

func (u *productUseCase) Update(ctx context.Context, newProduct Product, id string) (product Product, err error) {
	if product, err = u.productRepository.Update(ctx, newProduct, id); err != nil {
		return product, err
	}

	return product, nil
}
