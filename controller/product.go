package controller

import (
	// standard library packages
    "fmt"
	"database/sql"

	// third party packages
	_ "github.com/go-sql-driver/mysql"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/model"
)

type (
	ProductController struct {
		db *sql.DB
	}
	ProductTypeController struct{}
)

// ProductController apis
func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{db}
}

func (pc ProductController) Get(id int) model.Product {

	return model.Product{}
}

func (pc ProductController) Create(p model.Product) model.Product {
	return model.Product{}
}

func (pc ProductController) Update(p model.Product) model.Product {
	return model.Product{}
}

func (pc ProductController) Delete(id int) model.Product {
	return model.Product{}
}

func (pc ProductController) Query(s string) []model.Product {
	rows, err := pc.db.Query(s)
	if err != nil && rows == nil {
		return nil
	}
    defer pc.db.Close()
	products := make([]model.Product, 0)
    for rows.Next() {
        var products_id int
        var products_model string
        var type_id int
        var type_name string
        err = rows.Scan(&products_id, &products_model, &type_id, &type_name)
        if err != nil {
            panic(err)
        }
        fmt.Println(products_id)
        fmt.Println(products_model)
        fmt.Println(type_id)
        fmt.Println(type_name)
        products = append(products, model.Product{products_id, products_model, model.ProductType{type_id, type_name}})
    }
	return products
}
