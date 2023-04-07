package module_test

import (
	"context"
	"testing"

	"github.com/asaskevich/govalidator"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi3/module"
	"github.com/gusrylmubarok/test/tree/main/03-Go-Programming-Secure-Your-Go-Apps/Sesi3/module/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestInsert(t *testing.T) {
	mockAddedProduct := module.Product{
		ID:          "product-123",
		Title:       "A Product",
		Description: "Description A Product",
	}

	mockProductRepository := new(mocks.ProductRepository)
	productUseCase := module.NewProductUseCase(mockProductRepository)

	t.Run("should add product", func(t *testing.T) {
		tempMockProduct := module.Product{
			Title:       "A Product",
			Description: "Description A Product",
		}

		tempMockProduct.ID = "product-123"
		mockProductRepository.On("Insert", mock.Anything, mock.AnythingOfType("*module.Product")).Return(nil).Once()

		err := productUseCase.Insert(context.Background(), &tempMockProduct)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockProduct)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedProduct.ID, tempMockProduct.ID)
		assert.Equal(t, mockAddedProduct.Title, tempMockProduct.Title)
		assert.Equal(t, mockAddedProduct.Description, tempMockProduct.Description)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should error add product with empty title", func(t *testing.T) {
		tempMockProduct := module.Product{
			Title:       "",
			Description: "Description A Product",
		}

		tempMockProduct.ID = "product-123"
		mockProductRepository.On("Insert", mock.Anything, mock.AnythingOfType("*module.Product")).Return(nil).Once()

		err := productUseCase.Insert(context.Background(), &tempMockProduct)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockProduct)

		assert.Error(t, err)
		assert.Equal(t, mockAddedProduct.ID, tempMockProduct.ID)
		assert.NotEqual(t, mockAddedProduct.Title, tempMockProduct.Title)
		assert.Equal(t, mockAddedProduct.Description, tempMockProduct.Description)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should error add product with empty description", func(t *testing.T) {
		tempMockProduct := module.Product{
			Title:       "A Product",
			Description: "",
		}

		tempMockProduct.ID = "product-123"
		mockProductRepository.On("Insert", mock.Anything, mock.AnythingOfType("*module.Product")).Return(nil).Once()

		err := productUseCase.Insert(context.Background(), &tempMockProduct)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockProduct)

		assert.Error(t, err)
		assert.Equal(t, mockAddedProduct.ID, tempMockProduct.ID)
		assert.Equal(t, mockAddedProduct.Title, tempMockProduct.Title)
		assert.NotEqual(t, mockAddedProduct.Description, tempMockProduct.Description)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	mockUpdatedProduct := module.Product{
		ID:          "product-123",
		Title:       "A Product Updated",
		Description: "Description A Product Updated",
	}

	mockProductRepository := new(mocks.ProductRepository)
	productUseCase := module.NewProductUseCase(mockProductRepository)

	t.Run("should updated product", func(t *testing.T) {
		tempMockProductID := "product-123"
		tempMockUpdatedProduct := module.Product{
			Title:       "A Product Updated",
			Description: "Description A Product Updated",
		}

		mockProductRepository.On("Update", mock.Anything, mock.AnythingOfType("module.Product"), mock.AnythingOfType("string")).Return(mockUpdatedProduct, nil).Once()

		product, err := productUseCase.Update(context.Background(), tempMockUpdatedProduct, tempMockProductID)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockUpdatedProduct)

		assert.NoError(t, err)
		assert.Equal(t, product, mockUpdatedProduct)
		assert.Equal(t, mockUpdatedProduct.Title, tempMockUpdatedProduct.Title)
		assert.Equal(t, mockUpdatedProduct.Description, tempMockUpdatedProduct.Description)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should update product with empty message", func(t *testing.T) {
		tempMockProductID := "Product-123"
		tempMockUpdateProduct := module.Product{
			Title: "",
		}

		mockProductRepository.On("Update", mock.Anything, mock.AnythingOfType("module.Product"), mock.AnythingOfType("string")).Return(mockUpdatedProduct, nil).Once()

		product, err := productUseCase.Update(context.Background(), tempMockUpdateProduct, tempMockProductID)

		assert.NoError(t, err)

		tempMockUpdatedProduct := module.Product{
			ID:          "product-123",
			Title:       "A Product Updated",
			Description: "Description A Product Updated",
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateProduct)

		assert.Error(t, err)
		assert.Equal(t, product, tempMockUpdatedProduct)
		assert.Equal(t, mockUpdatedProduct.ID, tempMockUpdatedProduct.ID)
		assert.Equal(t, mockUpdatedProduct.Title, tempMockUpdatedProduct.Title)
		assert.Equal(t, mockUpdatedProduct.Description, tempMockUpdatedProduct.Description)
		mockProductRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockProduct := module.Product{
		ID:          "product-123",
		Title:       "A Product",
		Description: "Description A Product",
	}

	mockProductRepository := new(mocks.ProductRepository)
	productUseCase := module.NewProductUseCase(mockProductRepository)

	t.Run("should delete product", func(t *testing.T) {
		mockProductRepository.On("DeleteById", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := productUseCase.DeleteById(context.Background(), mockProduct.ID)

		assert.NoError(t, err)
		mockProductRepository.AssertExpectations(t)
	})

	t.Run("should not delete product wrong id", func(t *testing.T) {
		mockProductRepository.On("DeleteById", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := productUseCase.DeleteById(context.Background(), "Product-123")

		assert.NoError(t, err)
		mockProductRepository.AssertExpectations(t)
	})

}

func TestFindAll(t *testing.T) {
	mockProduct := module.Product{
		ID:          "product-123",
		Title:       "A Product",
		Description: "Description A Product",
	}

	mockProducts := make([]module.Product, 0)
	mockProducts = append(mockProducts, mockProduct)

	mockProductRepository := new(mocks.ProductRepository)
	productUseCase := module.NewProductUseCase(mockProductRepository)

	t.Run("find all products", func(t *testing.T) {
		mockProductRepository.On("FindAll", mock.Anything, mock.AnythingOfType("*[]module.Product"), mock.AnythingOfType("string")).Return(nil).Once()

		err := productUseCase.FindAll(context.Background(), &mockProducts)

		assert.NoError(t, err)
	})
}

func TestFindById(t *testing.T) {
	mockProduct := module.Product{
		ID:          "product-123",
		Title:       "A Product",
		Description: "Description A Product",
	}

	mockProductRepository := new(mocks.ProductRepository)
	productUseCase := module.NewProductUseCase(mockProductRepository)

	t.Run("should find by id", func(t *testing.T) {
		mockProductRepository.On("FindById", mock.Anything, mock.AnythingOfType("string")).Return(mockProduct, nil).Once()

		product, err := productUseCase.FindById(context.Background(), mockProduct.ID)

		assert.NoError(t, err)
		assert.Equal(t, mockProduct.ID, product.ID)
		mockProductRepository.AssertExpectations(t)
	})

}
