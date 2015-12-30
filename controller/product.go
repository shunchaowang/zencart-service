package controller

import (
	// standard library packages
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
	products := make([]model.Product, 10)
	return products
}
