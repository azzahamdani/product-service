package service

import "github.com/azzahamdani/product-service/repository"

type ProductService struct {
	repo *repository.ProductRepository
}

func NewProductService(repo *repository.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetProduct(productID string) (map[string]string, error) {
	product, err := s.repo.FindByID(productID)
	if err != nil {
		return nil, err
	}

	// Business logic can be applied here, for example, applying discounts
	// or calculating sales tax, etc.

	return product, nil
}
