package services

import (
	fetcher "my_application_name/materi/testing/main/dataaccess"
)

type Product struct {
	apiFetcher fetcher.ThirdPartyAPIInterface
}

func ProductHandler() Product {
	return Product{fetcher.ThirdPartyAPIHandler()}
}

type ProductInterFace interface {
	GetById(id int) fetcher.Product
	UpdateById(id int, payload fetcher.Product) interface{}
}

func (service *Product) GetById(id int) fetcher.Product {
	product := service.apiFetcher.GetProductByID(id)
	product.Name = "Prefix-" + product.Name
	return product
}

func (service *Product) UpdateById(id int, payload fetcher.Product) (interface{}) {
	service.apiFetcher.UpdateProductById(payload, id)
	return nil

}
