package controller

import (
	// standard library packages
	"database/sql"
	"fmt"

	// third party packages
	_ "github.com/go-sql-driver/mysql"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/model"
)

type (
	ProductController struct {
		db *sql.DB
	}

	CategoryController struct {
		db *sql.DB
	}
)

const productQuery = "SELECT p.products_id, p.products_model, p.products_quantity, p.products_image, p.products_price, pa.products_attributes_weight, pd.products_name, pd.products_description, c.categories_id, c.categories_image, cd.categories_name, cd.categories_description FROM products p JOIN products_attributes pa ON p.products_id = pa.products_id JOIN products_description pd ON p.products_id = pd.products_id JOIN products_to_categories ptc ON p.products_id = ptc.products_id JOIN categories c ON ptc.categories_id = c.categories_id JOIN categories_description cd ON cd.categories_id = c.categories_id"

const categoryQuery = "SELECT c.categories_id, c.categories_image, cd.categories_name, cd.categories_description FROM categories c JOIN categories_description cd ON cd.categories_id = c.categories_id"

// ProductController apis
func NewProductController(db *sql.DB) *ProductController {
	return &ProductController{db}
}

func (pc ProductController) Get(id int) model.Product {

	query := fmt.Sprintf("%s WHERE p.products_id = %d", productQuery, id)
	row := pc.db.QueryRow(query)
	defer pc.db.Close()
	var product_id int
	var product_model string
	var quantity int
	var image string
	var price float32
	var weight float32
	var name string
	var description string

	var category_id int
	var category_image string
	var category_name string
	var category_description string

	err := row.Scan(&product_id, &product_model, &quantity, &image, &price, &weight, &name, &description, &category_id, &category_image, &category_name, &category_description)
	if err != nil {
		panic(err)
	}
	product := model.Product{product_id, product_model, quantity, image, price, weight, name, description, model.Category{category_id, category_image, category_name, category_description}}
	return product
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

func (pc ProductController) GetAll() []model.Product {
	rows, err := pc.db.Query(productQuery)
	if err != nil && rows == nil {
		return nil
	}
	defer pc.db.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var id int
		var product_model string
		var quantity int
		var image string
		var price float32
		var weight float32
		var name string
		var description string

		var category_id int
		var category_image string
		var category_name string
		var category_description string

		err = rows.Scan(&id, &product_model, &quantity, &image, &price, &weight, &name, &description, &category_id, &category_image, &category_name, &category_description)
		if err != nil {
			panic(err)
		}
		products = append(products, model.Product{id, product_model, quantity, image, price, weight, name, description, model.Category{category_id, category_image, category_name, category_description}})
	}
	return products
}

func (pc ProductController) GetPage(limit, offset int) []model.Product {
	s := fmt.Sprintf("%s LIMIT %d OFFSET %d", productQuery, limit, offset)
	rows, err := pc.db.Query(s)
	if err != nil && rows == nil {
		return nil
	}
	defer pc.db.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var id int
		var product_model string
		var quantity int
		var image string
		var price float32
		var weight float32
		var name string
		var description string

		var category_id int
		var category_image string
		var category_name string
		var category_description string

		err = rows.Scan(&id, &product_model, &quantity, &image, &price, &weight, &name, &description, &category_id, &category_image, &category_name, &category_description)
		if err != nil {
			panic(err)
		}
		products = append(products, model.Product{id, product_model, quantity, image, price, weight, name, description, model.Category{category_id, category_image, category_name, category_description}})
	}
	return products
}

/*
 * Query supports up to 3 args, query criteria, pagination.
 * 1 arg: must be string, search product by name or description
 * 2 args: must be 2 ints, the first one is page size, the second one is page number
 * 3 args: must be string, int, int. Combine up 2.
 */
func (pc ProductController) Query(args ...interface{}) []model.Product {
	// Get any parameters passed out of the args variable
	var query string
	switch len(args) {
	case 0: // query all
		query = productQuery
	case 1: // query criteria
		criteria, ok := args[0].(string)
		if !ok {
			panic("1st parameter not type string.")
		}
		query = fmt.Sprintf("%s WHERE pd.products_name LIKE '%%s%%'", productQuery, criteria, criteria)
	case 2: // paged
		limit, ok := args[0].(int)
		if !ok {
			panic("1st parameter not type int.")
		}
		offset, ok := args[1].(int)
		if !ok {
			panic("2nd parameter not type int.")
		}
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", productQuery, limit, offset)
	case 3: // paged query criteria
		criteria, ok := args[0].(string)
		if !ok {
			panic("1st parameter not type string.")
		}
		limit, ok := args[1].(int)
		if !ok {
			panic("2nd parameter not type int.")
		}
		offset, ok := args[2].(int)
		if !ok {
			panic("3rd parameter not type int.")
		}
		query = fmt.Sprintf("%s WHERE pd.products_name LIKE '%%%s%%' OR pd.products_description LIKE '%%%s%%' LIMIT %d OFFSET %d", productQuery, criteria, criteria, limit, offset)
	default:
		query = productQuery
	}
	rows, err := pc.db.Query(query)
	if err != nil && rows == nil {
		return nil
	}
	defer pc.db.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var id int
		var product_model string
		var quantity int
		var image string
		var price float32
		var weight float32
		var name string
		var description string

		var category_id int
		var category_image string
		var category_name string
		var category_description string

		err = rows.Scan(&id, &product_model, &quantity, &image, &price, &weight, &name, &description, &category_id, &category_image, &category_name, &category_description)
		if err != nil {
			panic(err)
		}
		products = append(products, model.Product{id, product_model, quantity, image, price, weight, name, description, model.Category{category_id, category_image, category_name, category_description}})
	}
	return products
}

/*
 * Get products by category.
 * Support 1 argument or 3 argument, 1st is category id, 2nd limit and 3rd offset are for pagination.
 */
func (pc ProductController) GetProductsByCategory(category_id int, args ...interface{}) []model.Product {
	var query string
	if len(args) == 2 {
		limit, ok := args[0].(int)
		if !ok {
			panic("1st parameter not type int.")
		}
		offset, ok := args[1].(int)
		if !ok {
			panic("2nd parameter not type int.")
		}
		query = fmt.Sprintf("%s WHERE c.categories_id = %d LIMIT %d OFFSET %d", productQuery, category_id, limit, offset)
	} else {
		query = fmt.Sprintf("%s WHERE c.categories_id = %d", productQuery, category_id)
	}
	rows, err := pc.db.Query(query)
	if err != nil && rows == nil {
		return nil
	}
	defer pc.db.Close()
	products := make([]model.Product, 0)
	for rows.Next() {
		var id int
		var product_model string
		var quantity int
		var image string
		var price float32
		var weight float32
		var name string
		var description string

		var category_id int
		var category_image string
		var category_name string
		var category_description string

		err = rows.Scan(&id, &product_model, &quantity, &image, &price, &weight, &name, &description, &category_id, &category_image, &category_name, &category_description)
		if err != nil {
			panic(err)
		}
		products = append(products, model.Product{id, product_model, quantity, image, price, weight, name, description, model.Category{category_id, category_image, category_name, category_description}})
	}
	return products
}

// Category Controller apis
func NewCategoryController(db *sql.DB) *CategoryController {
	return &CategoryController{db}
}

func (cc CategoryController) Get(id int) model.Category {
	query := fmt.Sprintf("%s WHERE c.categories_id = %d", categoryQuery, id)
	row := cc.db.QueryRow(query)
	defer cc.db.Close()
	var category_id int
	var image string
	var name string
	var description string

	err := row.Scan(&category_id, &image, &name, &description)

	if err != nil {
		panic(err)
	}
	category := model.Category{id, image, name, description}
	return category
}

func (cc CategoryController) GetAll() []model.Category {
	rows, err := cc.db.Query(categoryQuery)
	if err != nil || rows == nil {
		return nil
	}
	defer cc.db.Close()
	categories := make([]model.Category, 0)
	for rows.Next() {
		var id int
		var image string
		var name string
		var description string

		err = rows.Scan(&id, &image, &name, &description)

		if err != nil {
			panic(err)
		}
		categories = append(categories, model.Category{id, image, name, description})
	}
	return categories
}

/*
 * Query supports up to 3 args, query criteria, pagination.
 * 1 arg: must be string, search product by name or description
 * 2 args: must be 2 ints, the first one is page size, the second one is page number
 * 3 args: must be string, int, int. Combine up 2.
 */
func (cc CategoryController) Query(args ...interface{}) []model.Category {
	// Get any parameters passed out of the args variable
	var query string
	switch len(args) {
	case 0: // query all
		query = categoryQuery
	case 1: // query criteria
		criteria, ok := args[0].(string)
		if !ok {
			panic("1st parameter not type string.")
		}
		query = fmt.Sprintf("%s WHERE cd.categories_name LIKE '%%s%%'", categoryQuery, criteria, criteria)
	case 2: // paged
		limit, ok := args[0].(int)
		if !ok {
			panic("1st parameter not type int.")
		}
		offset, ok := args[1].(int)
		if !ok {
			panic("2nd parameter not type int.")
		}
		query = fmt.Sprintf("%s LIMIT %d OFFSET %d", categoryQuery, limit, offset)
	case 3: // paged query criteria
		criteria, ok := args[0].(string)
		if !ok {
			panic("1st parameter not type string.")
		}
		limit, ok := args[1].(int)
		if !ok {
			panic("2nd parameter not type int.")
		}
		offset, ok := args[2].(int)
		if !ok {
			panic("3rd parameter not type int.")
		}
		query = fmt.Sprintf("%s WHERE cd.categories_name LIKE '%%%s%%' LIMIT %d OFFSET %d", categoryQuery, criteria, criteria, limit, offset)
	default:
		query = categoryQuery
	}
	rows, err := cc.db.Query(query)
	if err != nil && rows == nil {
		return nil
	}
	defer cc.db.Close()
	categories := make([]model.Category, 0)
	for rows.Next() {
		var id int
		var image string
		var name string
		var description string

		err = rows.Scan(&id, &image, &name, &description)
		if err != nil {
			panic(err)
		}
		categories = append(categories, model.Category{id, image, name, description})
	}
	return categories
}
