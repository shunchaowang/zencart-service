package controller

import (
	// standard library packages
	"database/sql"
	"fmt"
    "crypto/md5"

	// third party packages
	_ "github.com/go-sql-driver/mysql"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/model"
)

type CustomerController struct {
    db *sql.DB
}
// salt const for password generation
const salt = "zc"

const customerQuery = "SELECT p.products_id, p.products_model, p.products_quantity, p.products_image, p.products_price, pa.products_attributes_weight, pd.products_name, pd.products_description, c.categories_id, c.categories_image, cd.categories_name, cd.categories_description FROM products p JOIN products_attributes pa ON p.products_id = pa.products_id JOIN products_description pd ON p.products_id = pd.products_id JOIN products_to_categories ptc ON p.products_id = ptc.products_id JOIN categories c ON ptc.categories_id = c.categories_id JOIN categories_description cd ON cd.categories_id = c.categories_id"

const createCustomerSql = "INSERT INTO customers (customers_gender, customers_firstname, customers_lastname, customers_dob, customers_email_address, customers_nick,) VALUES ()"

// CustomerController apis
func NewCustomerController(db *sql.DB) *CustomerController {
    return &CustomerController{db}
}

// get customer
func (cc CustomerController) Get(id int) model.Customer {
    return model.Customer{}
}
// query customer by email
func (cc CustomerController) QueryByEmail(email string) model.Customer {
    return model.Customer{}
}






// query customer
// create customer
func (cc CustomerController) Create(c model.Customer) model.Customer {
    return model.Customer{}
}
// update customer
// delete customer
