package dataaccess

type ThirdPartyAPI struct {
}

type Product struct{
	ID int
	Name string
	Price float32
	IsDeleted bool
}

type ThirdPartyAPIInterface interface {
	GetProductByID(int) Product
	UpdateProductById(Product,int) bool
}

func ThirdPartyAPIHandler() *ThirdPartyAPI{
	return &ThirdPartyAPI{}
}

func (apiFetcher *ThirdPartyAPI) GetProductByID(id int) Product {
	response := Product{} // replace with all magic stuff here	
	return response
}

func (apiFetcher *ThirdPartyAPI) UpdateProductById(payload Product, id int) bool {
	response := false // replace with all magic stuff here	
	return response
}
