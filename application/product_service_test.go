package application_test

import (
	"github.com/othiagosilva/arquitetura-hexagonal/application"
	mock_application "github.com/othiagosilva/arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestProductService_Get(t *testing.T) {
	
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	product := mock_application.NewMockProductInterface(mockCtrl)
	persistence := mock_application.NewMockProductPersistenceInterface(mockCtrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}
	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}

func TestProductService_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	product := mock_application.NewMockProductInterface(mockCtrl)
	persistence := mock_application.NewMockProductPersistenceInterface(mockCtrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Create("Product 1", 10)
	require.Nil(t, err)
	require.Equal(t, product, result)
}

func TestProductService_EnableDisable(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	product := mock_application.NewMockProductInterface(mockCtrl)
	product.EXPECT().Enable().Return(nil)
	product.EXPECT().Disable().Return(nil)

	persistence := mock_application.NewMockProductPersistenceInterface(mockCtrl)
	persistence.EXPECT().Save(gomock.Any()).Return(product, nil).AnyTimes()
	service := application.ProductService{Persistence: persistence}

	result, err := service.Enable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)

	result, err = service.Disable(product)
	require.Nil(t, err)
	require.Equal(t, product, result)
}