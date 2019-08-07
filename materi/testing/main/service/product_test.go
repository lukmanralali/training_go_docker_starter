package services

import (
	"testing"
	"fmt"
	fetcher "my_application_name/materi/testing/main/dataaccess"
)
import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ThirdPartyAPIMock struct {
	mock.Mock
}
func (mockSample *ThirdPartyAPIMock) GetProductByID(id int) fetcher.Product {
	mockSample.Called(id)
	resultMock := fetcher.Product{}
	resultMock.ID = id
	resultMock.Name = "test product name"
	resultMock.Price = float32(10000)
	resultMock.IsDeleted = false
	return resultMock
}

func (mockSample *ThirdPartyAPIMock) UpdateProductById(payload fetcher.Product, id int) bool {
	mockSample.Called(payload, id)
	return true
}

// TEST
func TestGetProductByID(t *testing.T) {
	// t.Parallel()

	dataFetcher := ThirdPartyAPIMock{}
	testId := 1

	dataFetcher.On("GetProductByID", testId).Return(fetcher.Product{})

	productService := Product{&dataFetcher}
	resultFuncService := productService.GetById(testId)

	assert.Equal(t, testId, resultFuncService.ID, "It should be same ID")
	assert.Equal(t, "Prefix-test product name", resultFuncService.Name, "It should be same Name")
	assert.Equal(t, float32(10000), resultFuncService.Price, "It should be same Price")
	assert.Equal(t, false, resultFuncService.IsDeleted, "It should be false")
}

func TestUpdateProductByID(t *testing.T) {
	// t.Parallel()

	dataFetcher := ThirdPartyAPIMock{}
	testId := 1

	samplePayload := fetcher.Product{
		ID: 1,
		Name: "Update Data",
		Price: float32(100),
		IsDeleted: true,
	}

	dataFetcher.On("UpdateProductById", samplePayload, testId).Return(true)

	productService := Product{&dataFetcher}
	resultFuncService := productService.UpdateById(testId, samplePayload)
	fmt.Println(resultFuncService)
	assert.Equal(t, nil, resultFuncService, "It should be nil because not implemented yet")
}
