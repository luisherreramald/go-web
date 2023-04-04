package products

import (
	"api/api/internal/domain"
	"errors"
)

type Repository interface {
	GetById(id int) (product *domain.Product, err error)
	Create(product *domain.Product) (lastId int, err error)
	Update(product *domain.Product, id int) (err error)
	Delete(id int) (err error)
}

var (
	ErrRepoNotUnique = errors.New("The code value is not unique")
	ErrRepoNotFound = errors.New("Product Not Found")
	ErrRepoInternalError = errors.New("Internal Error")
)

type Service interface {
	GetById(id int) (product *domain.Product, err error)
	Create(product *domain.Product) (err error)
	Update(product *domain.Product, id int) (err error)
	Delete(id int) (err error)
}

var (
	ErrServiceNotUnique = errors.New("The code value is not unique")
	ErrServiceNotFound = errors.New("Product Not Found")
	ErrServiceInternalError = errors.New("Internal Server Error")
)