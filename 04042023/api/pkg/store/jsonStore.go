package store

import (
	"api/api/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

const fileName = "products.json"

var (
	ErrStoreNotFound = errors.New("Product Not Found")
	ErrStoreNotUnique = errors.New("Code value is not unique")
)


func LoadProducts() (err error) {

	file, err := os.Open(fileName)

	if err != nil {
		return 
	}

	_, err = ioutil.ReadAll(file)

	return 
}

func ProductJsonFileToArray() (productsArray []domain.Product, err error) {
	fileReader, err := os.OpenFile(fileName, os.O_RDONLY, 0777)
	if err != nil {
		return 
	}
	
	products, err := ioutil.ReadAll(fileReader)
	
	if err != nil {
		return
	}
	
	json.Unmarshal(products, &productsArray)	

	fileReader.Close()

	return
}

func SearchProduct(id int) (product domain.Product, err error) {
	productsArray, err := ProductJsonFileToArray() 
	
	if err != nil {
		return
	}

	for _, product := range productsArray {
		if product.Id == id {
			return product, nil
		}
	}
	
	return domain.Product{}, ErrStoreNotFound
}

func ValidateCodeValue(product *domain.Product) (err error) {
	productsArray, err := ProductJsonFileToArray() 
	
	if err != nil {
		return
	}
	
	for _, productStore := range productsArray {
		if product.CodeValue == productStore.CodeValue  {
			return ErrStoreNotUnique
		}
	}

	return 
}

func ValidateCodeValueUpdate(product *domain.Product, id int) (err error) {
	productsArray, err := ProductJsonFileToArray() 
	
	if err != nil {
		return
	}
	
	for _, productStore := range productsArray {
		fmt.Printf("id: %d, idInt: %d", id, productStore.Id)
		if product.CodeValue == productStore.CodeValue && productStore.Id != id  {
			return ErrStoreNotUnique
		}
	}

	return 
}

func GetAllProducts() (products []domain.Product, err error) {

	products, err = ProductJsonFileToArray()

	if err != nil {
		return
	}

	return 
}

func SaveStructToFile(products []domain.Product) (err error) {
	data, err:= json.MarshalIndent(products, "", " ")

	if err != nil {
		return 
	}

	err = ioutil.WriteFile(fileName, data, 0777)

	return 
}

func CreateProduct(product *domain.Product) (err error) {
	productsArray, err := ProductJsonFileToArray() 
	
	if err != nil {
		return
	}

	newId := productsArray[len(productsArray)-1].Id + 1

	product.Id = newId
	productsArray = append(productsArray, *product)
	data, err:= json.MarshalIndent(productsArray, "", " ")

	if err != nil {
		return 
	}

	err = ioutil.WriteFile(fileName, data, 0777)

 	return
}

func Update(product *domain.Product, id int) (err error) {
	productsArray, err := ProductJsonFileToArray()
	
	if err != nil {
		return 
	}

	for index, productIterator := range productsArray {
		if productIterator.Id == id {
			productsArray[index] = *product
			err = SaveStructToFile(productsArray)
			return 
		}
	}

	return ErrStoreNotFound
}

func Delete(id int) (err error) {
	productsArray, err := ProductJsonFileToArray()
	
	if err != nil {
		return 
	}

	for index, productIterator := range productsArray {
		if productIterator.Id == id {
			productsArray = append(productsArray[:index], productsArray[index+1:]...)
			err = SaveStructToFile(productsArray)
			return 
		}
	}

	return ErrStoreNotFound
}