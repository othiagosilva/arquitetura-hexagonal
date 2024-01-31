package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	db2 "github.com/othiagosilva/arquitetura-hexagonal/adapters/db"
	"github.com/othiagosilva/arquitetura-hexagonal/application"
)

func main() {
	db, _ := sql.Open("sqlite3", "sqlite.db")
	productDbAdapter := db2.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ :=productService.Create("Product 1", 10.0)

	productService.Enable(product)	
}