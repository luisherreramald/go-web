package products

import (
	"api/api/internal/domain"
	"errors"
)

func NewService(rp Repository) Service {
	return &service{rp: rp}
}

type service struct {
	rp Repository
}

func (productService *service) GetById(id int) (product domain.Product, err error) {
	product, err = productService.rp.GetById(id)
	if err != nil {
		if errors.Is(err, ErrRepoNotFound) {
			err = ErrServiceNotFound
			return 
		}
	}

	return
}

func (productService *service) Create(product *domain.Product) (err error) {
	err = product.Validate()
	if err != nil {
		return 
	}

	err = productService.rp.Create(product)

	if err != nil {
		if errors.Is(err, ErrRepoNotUnique) {
			err = ErrServiceNotUnique
			return
		}
		return ErrServiceInternalError
	}
	return 
}

func (productService *service) Update(product *domain.Product, id int) (err error) {
	err = product.Validate()
	if err != nil {
		return 
	}

	err = productService.rp.Update(product, id)
	if err != nil {
		if errors.Is(err, ErrRepoNotFound) {
			err = ErrServiceNotFound
			return
		}
		
		if errors.Is(err, ErrRepoNotUnique) {
			return ErrServiceNotUnique
		}

		return ErrServiceInternalError
	}

	product.Id = id

	return 
}

func (productService *service) Delete(id int) (err error) {

	err = productService.rp.Delete(id)
	if err != nil {
		if errors.Is(err, ErrRepoNotFound) {
			err = ErrServiceNotFound
			return
		}
		return ErrServiceInternalError
	}	
	
	return 
}