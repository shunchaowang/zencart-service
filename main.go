package main

import (
	// standard library packages
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// third party packages
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"

	// project scope packages
	//"github.com/shunchaowang/smartcart-service/model"
	"github.com/shunchaowang/zencart-service/controller"
)

func main() {
	router := httprouter.New()

	router.GET("/product/query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		pc := controller.NewProductController(getMysqlDB())
		products := pc.Query()
		psj, _ := json.Marshal(products)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", psj)
	})

	router.GET("/product/get", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		pc := controller.NewProductController(getMysqlDB())
		idParam := r.FormValue("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			panic("id not converted to type int.")
		}
		product := pc.Get(id)
		pj, _ := json.Marshal(product)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", pj)
	})

	router.GET("/category/products", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		pc := controller.NewProductController(getMysqlDB())
		idParam := r.FormValue("cid")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			panic("id not converted to type int.")
		}
		products := pc.GetProductsByCategory(id)
		psj, _ := json.Marshal(products)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", psj)
	})

	router.GET("/category/get", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cc := controller.NewCategoryController(getMysqlDB())
		idParam := r.FormValue("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			panic("id not converted to type int.")
		}
		category := cc.Get(id)
		cj, _ := json.Marshal(category)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", cj)
	})

	router.GET("/category/query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cc := controller.NewCategoryController(getMysqlDB())
		categories := cc.Query()
		csj, _ := json.Marshal(categories)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", csj)
	})

	//TODO: DELETE
	router.GET("/pagedProducts", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		pc := controller.NewProductController(getMysqlDB())
		//products := pc.Query("SELECT products_id, products_model, type_id, type_name FROM products, product_types WHERE products.products_type = product_types.type_id LIMIT 10 OFFSET 5")
		products := pc.Query(10, 5)
		psj, _ := json.Marshal(products)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", psj)
		/*for _, product := range products {

		  }*/
	})

	//TODO: DELETE
	router.GET("/allCategories", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		cc := controller.NewCategoryController(getMysqlDB())
		//products := pc.Query("SELECT products_id, products_model, type_id, type_name FROM products, product_types WHERE products.products_type = product_types.type_id LIMIT 10 OFFSET 5")
		categories := cc.GetAll()
		csj, _ := json.Marshal(categories)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", csj)
		/*for _, product := range products {

		  }*/
	})

	//http.ListenAndServe("localhost:8080", router)
	http.ListenAndServeTLS(":7443", "/Users/swang/Developer/go/cert/gocert.pem", "/Users/swang/Developer/go/cert/gokey.pem", router)
}

func getMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "zencart:zencart@tcp(localhost:3306)/zencart?charset=utf8")

	if err != nil {
		panic(err)
	}

	return db
}
