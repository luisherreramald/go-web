package products

import (
	"api/api/internal/domain"
)

func NewRepositoryLocal(db []*domain.Product, lastId int) Repository {
	return &repositoryLocal{db: db, lastId: lastId}
}

type repositoryLocal struct {
	db []*domain.Product
	lastId int
}

func (rp *repositoryLocal) GetById(id int) (product *domain.Product, err error) {
	for _, productArray := range rp.db {
		if productArray.Id == id {
			product = productArray
			return 
		}
	}
	err = ErrRepoNotFound

	return 
}

func (rp *repositoryLocal) Create(product *domain.Product) (lastId int, err error) {

	for _, productArray := range rp.db {
		if product.CodeValue == productArray.CodeValue {
			return 0, ErrRepoNotUnique
		}
	}

	rp.lastId++
	product.Id = rp.lastId

	rp.db = append(rp.db, product)
	lastId = rp.lastId

	return 
}

func (rp *repositoryLocal) Update(product *domain.Product, id int) (err error) {
	for index, productArray := range rp.db {
		if productArray.Id == id {
			rp.db[index] = product
			return 
		}
	}
	
	err = ErrRepoNotFound
	
	return 
}

func (rp *repositoryLocal) Delete(id int) (err error) {

	for index, product := range rp.db {
		if product.Id == id {
			rp.db = append(rp.db[:index], rp.db[index+1:]...)

			return 
		}
	}

	err = ErrRepoNotFound

	return
}