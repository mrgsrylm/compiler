package module

import (
	"context"
	"fmt"
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

func (u *productUseCase) DeleteById(ctx context.Context, id string) (err error) {
	if err = u.productRepository.DeleteById(ctx, id); err != nil {
		return err
	}

	return
}

func (u *productUseCase) FindAll(ctx context.Context, products *[]Product) (err error) {
	if err = u.productRepository.FindAll(ctx, products); err != nil {
		return err
	}
	return
}

func (u *productUseCase) FindById(ctx context.Context, id string) (product Product, err error) {
	if product, err = u.productRepository.FindById(ctx, id); err != nil {
		return product, err
	}
	fmt.Println(product)

	return product, nil
}
