package application_test

import (
	"github.com/othiagosilva/arquitetura-hexagonal/application"
	"testing"
	"github.com/stretchr/testify/require"
)

func TestProduct_Enable(t *testing.T) {
	product := application.Product{}
	product.Name = "Hello"
	product.Status = application.DISABLED
	product.Price = 10
}