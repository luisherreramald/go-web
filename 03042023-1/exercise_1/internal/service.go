package internal

import (
	"errors"
	"time"
)

var (
	ErrProductExist = errors.New("The code value is registered, insert a new code value")
	ErrProductNotFound = errors.New("Product Not Found")
)

type Product struct {
	Id int `json:"id"`
	Name string `json:"name" validate:"required"`
	Quantity int `json:"quantity" validate:"number, required"`
	CodeValue string `json:"code_value" validate:"required"`
	IsPublished bool `json:"is_published"`
	Expiration time.Time `json:"expiration" validate:"required, datetime"`
	Price float64 `json:"price" validate:"number,required"`
}

func validateCodeValue(product *Product, products []*Product) bool {
	for _, productArray := range products {
		if product.CodeValue == productArray.CodeValue {
			return false
		}
	}

	return true
}

func NewServiceProduct(db []*Product, lastID int) *ServiceProducts {
	return &ServiceProducts{
		db: db,
		lastID: lastID,
	}
}

type ServiceProducts struct {
	db   []*Product
	lastID int
}


func (serviceProduct *ServiceProducts) Save(product *Product) (productResult *Product, err error) {
	product.Id = serviceProduct.lastID



	if validateCodeValue(product, serviceProduct.db) == false {
		return nil, ErrProductExist

	}

	serviceProduct.db = append(serviceProduct.db, product)

	

	serviceProduct.lastID++

	return product, nil
}

func (serviceProduct *ServiceProducts) GetProductById(id int) (product *Product, err error) {
	for _, productArray := range serviceProduct.db {
		if productArray.Id == id {
			return productArray, nil
		}
	}

	return nil, ErrProductNotFound
}


