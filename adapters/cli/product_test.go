package cli_test

import (
	"fmt"
	"github.com/othiagosilva/arquitetura-hexagonal/adapters/cli"
	mock_application "github.com/othiagosilva/arquitetura-hexagonal/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)
func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	
	productName := "Product 1"
	productPrice := 25.99
	productStatus := "enabled"
	productId := "abc"

	mockApp := mock_application.NewMockProductInterface(ctrl)
	mockApp.EXPECT().GetID().Return(productId).AnyTimes()
	mockApp.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	mockApp.EXPECT().GetName().Return(productName).AnyTimes()
	mockApp.EXPECT().GetPrice().Return(productPrice).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(mockApp, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(mockApp, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(mockApp, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(mockApp, nil).AnyTimes()

	resultExpected := fmt.Sprintf("Product ID %s with the name %s has been created with the price %f and status %s", productId, productName, productPrice, productStatus)
	result, err := cli.Run(service, "create", "", productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been enabled", productName)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product %s has been disabled", productName)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)

	resultExpected = fmt.Sprintf("Product ID %s with the name %s, price %f and status %s", productId, productName, productPrice, productStatus)
	result, err = cli.Run(service, "get", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, resultExpected, result)
}