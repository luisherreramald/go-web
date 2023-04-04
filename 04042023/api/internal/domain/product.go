package domain

import "errors"
import "time"

type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Quantity int `json:"quantity"`
	CodeValue string `json:"code_value"`
	IsPublished bool `json:"is_published"`
	Expiration string `json:"expiration"`
	Price float64 `json:"price"`
}


func (product *Product ) Validate() (err error) {

	_, err = time.Parse("01/02/2006", product.Expiration)

	if err != nil {
		err = errors.New("Expiration date is not valid")
		return 
	}

	return nil
}