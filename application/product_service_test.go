package application_test

import (
	"testing"
	"github.com/stretchr/testify/require"
	"github.com/golang/mock/gomock"
)

func TestProductService_Get(t *testing.T) {
	
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	product : mock_application.NewMockProductInterface(mockCtrl)
	persistence := mock_application.NewMockProductPersistenceInterface(mockCtrl)
	persistence.EXPECT().Get(gomock.Any()).Return(product, nil).AnyTimes()

	service := application.ProductService{Persistence: persistence}
	result, err := service.Get("abc")
	require.Nil(t, err)
	require.Equal(t, product, result)

}