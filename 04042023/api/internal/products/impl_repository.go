package products

import (
	"api/api/internal/domain"
	"api/api/pkg/store"
	"errors"
)

func NewRepositoryLocal() Repository {
	return &repositoryLocal{}
}

type repositoryLocal struct {

}

func (rp *repositoryLocal) GetById(id int) (product domain.Product, err error) {
	product, err = store.SearchProduct(id)

	if err != nil  {
		if errors.Is(err, store.ErrStoreNotFound){
			err = ErrRepoNotFound
			return 
		}
		return domain.Product{}, ErrServiceInternalError
	}

	return product, nil
}

func (rp *repositoryLocal) Create(product *domain.Product) (err error) {
	err = store.ValidateCodeValue(product)
	
	if err != nil {
		if errors.Is(err, store.ErrStoreNotUnique) {
			err = ErrRepoNotUnique
		}
		return 
	}

	err = store.CreateProduct(product)
	return 
}

func (rp *repositoryLocal) Update(product *domain.Product, id int) (err error) {
	err = store.ValidateCodeValueUpdate(product, id)
	
	if err != nil {

		if errors.Is(err, store.ErrStoreNotUnique) {
			err = ErrRepoNotUnique
		}
		return 
	}

	err = store.Update(product, id)

	if err != nil {
		if errors.Is(err, store.ErrStoreNotFound) {
			err = ErrRepoNotFound
		}
		return
	}

	return 
}

func (rp *repositoryLocal) GetAllProducts() (products []domain.Product, err error) {
	products, err = store.GetAllProducts()

	return
}

func (rp *repositoryLocal) Delete(id int) (err error) {
	err = store.Delete(id)

	if err != nil {
		if errors.Is(err, store.ErrStoreNotFound) {
			err = ErrRepoNotFound
		}

		return 
	}
	return
}