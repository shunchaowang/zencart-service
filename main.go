package main

import (
	// standard library packages
	"fmt"
	"encoding/json"
	"net/http"
	"database/sql"

	// third party packages
    "github.com/julienschmidt/httprouter"
    _ "github.com/go-sql-driver/mysql"

    // project scope packages
    //"github.com/shunchaowang/smartcart-service/model"
    "github.com/shunchaowang/zencart-service/controller"
)

func main() {
    router := httprouter.New();

    router.GET("/query", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
        pc := controller.NewProductController(getMysqlDB())
        products := pc.Query("SELECT products_id, products_model FROM products")
        psj, _ := json.Marshal(products)

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(200)
        fmt.Fprintf(w, "%s", psj)
        /*for _, product := range products {
        
        }*/
    })

    http.ListenAndServe("localhost:8080", router)
}

func getMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "zencart:zencart@tcp(localhost:3306)/zencart?charset=utf8")

	if err != nil {
		panic(err)
	}

	return db
}
