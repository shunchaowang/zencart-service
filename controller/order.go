package controller

import (
	// standard library packages
	"database/sql"
	"fmt"
	"time"

	// third party packages
	_ "github.com/go-sql-driver/mysql"

	// project scope packages
	"github.com/shunchaowang/smartcart-service/model"
)
type (
	OrderController struct {
		db *sql.DB
	}
)


const userOrdersQuery = `
 	SELECT o.orders_id, ot.value as order_total, s.orders_status_id, s.orders_status_name, 
		o.date_purchased, o.last_modified, o.delivery_name,
		o.delivery_company, o.delivery_street_address, o.delivery_suburb,
		o.delivery_city, o.delivery_postcode, o.delivery_state, o.delivery_country,
		o.billing_name, o.billing_company,
		o.billing_street_address, o.billing_suburb, o.billing_city, o.billing_postcode,
		o.billing_state, o.billing_country
	FROM orders o 
	JOIN orders_total ot ON o.orders_id = ot.orders_id 
	JOIN orders_status s ON o.orders_status = s.orders_status_id`
	
const orderQuery = `
	select cc_cvv, customers_name, customers_company, customers_street_address,
		customers_suburb, customers_city, customers_postcode, customers_id,
		customers_state, customers_country, customers_telephone,
		customers_email_address, customers_address_format_id, delivery_name,
		delivery_company, delivery_street_address, delivery_suburb,
		delivery_city, delivery_postcode, delivery_state, delivery_country,
		delivery_address_format_id, billing_name, billing_company,
		billing_street_address, billing_suburb, billing_city, billing_postcode,
		billing_state, billing_country, billing_address_format_id,
		coupon_code, payment_method, payment_module_code, shipping_method, shipping_module_code,
		cc_type, cc_owner, cc_number, cc_expires, currency,
		currency_value, date_purchased, orders_status, last_modified,
		order_total, order_tax, ip_address
	from orders
	where orders_id = `
	
const orderProductQuery = `
	select orders_products_id, products_id, products_name, products_model,
		products_price, products_tax, products_quantity,
		final_price, onetime_charges,
		product_is_free from orders_products`
	
	
	
func (oc OrderController) GetCustomerOrders(customerId int) []model.Order {

	orderSql := fmt.Sprintf("%s WHERE o.customers_id = %d AND ot.class = 'ot_total' AND s.language_id = 1 ORDER BY orders_id DESC", userOrdersQuery, customerId)

	rows, err := oc.db.Query(orderSql)
	if err != nil && rows == nil {
		return nil
	}
	defer oc.db.Close()
	orders := make([]model.Order, 0)
	for rows.Next() {
		var orders_id int
		var order_total float32
		var orders_status_id int
		var orders_status_name string
		var date_purchased time.Time
		var last_modified time.Time
		
		var delivery_name string
		var delivery_company string
		var delivery_street_address string
		var delivery_suburb string
		var delivery_city string
		var delivery_postcode string
		var delivery_state string
		var delivery_country string
		var billing_name string
		var billing_company string
		var billing_street_address string
		var billing_suburb string
		var billing_city string
		var billing_postcode string
		var billing_state string
		var billing_country string
		
		err = rows.Scan(&orders_id, 
			&order_total, &orders_status_id, &orders_status_name,
			&date_purchased, &last_modified, &delivery_name,
			&delivery_company, &delivery_street_address, &delivery_suburb,
			&delivery_city, &delivery_postcode, &delivery_state, &delivery_country,
			&billing_name, &billing_company, &billing_street_address, 
			&billing_suburb, &billing_city, &billing_postcode,
			&billing_state, &billing_country)
		if err != nil {
			panic(err)
		}
		
		var shippingAddress = model.Address{delivery_company, "", "", delivery_name,
			delivery_street_address, delivery_suburb, delivery_city, delivery_state,
			delivery_postcode, delivery_country}
		var billingAddress = model.Address{billing_company, "", "", billing_name,
			billing_street_address, billing_suburb, billing_city, billing_state,
			billing_postcode, billing_country}
		
		
		products := make([]model.Product, 0)
		productSql := fmt.Sprintf("%s WHERE orders_id = %d order by orders_products_id", orderProductQuery, orders_id)
		productRows, err := oc.db.Query(productSql)
		if err != nil && productRows == nil {
			
		} else {
			var orders_products_id int
			var products_id int
			var products_name string
			var products_model string
			var products_price float32
			var products_tax float32
			var products_quantity int
			var final_price float32
			var onetime_charges float32
			var product_is_free string
			
			var image string
			var weight float32
			var description string
			
			err = rows.Scan(&orders_products_id, &products_id, &products_name, &products_model, 
				&products_price, &products_tax, &products_quantity, &final_price, &onetime_charges, 
				&product_is_free)
				
			if err != nil {
				panic(err)
			}
			products = append(products, model.Product{products_id, products_model, 
				products_quantity, image, final_price, weight, products_name, 
				description, model.Category{0, "", "", ""}})
		}
		
		
		orders = append(orders, model.Order{orders_id, order_total, "", 
			date_purchased, last_modified,
			model.OrderStatus{orders_status_id, orders_status_name}, 
			products, shippingAddress, billingAddress})
	}
	return orders
	
}


