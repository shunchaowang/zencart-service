package main

import (
	// standard library packages
	"database/sql"
	"net/http"

	// third party packages
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	// project scope packages
	"github.com/shunchaowang/zencart-service/controller"
)

func main() {

	// instantiate a new router
	r := httprouter.New()

	// get a ProductController instance
	pc := controller.NewProductController(getSession())

	// get a product
	r.GET("/product/:id", pc.Get)

	// create a product
	r.POST("/product", pc.Create)

	// delete a product
	r.DELETE("/product/:id", pc.Delete)

	// fire up the server
	http.ListenAndServe("localhost:8080", r)
}

// Create a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// Connect to our local mongo
	s, err := mgo.Dial("mongodb://localhost")

	// Check if connection fails, is mongo running?
	if err != nil {
		panic(err)
	}

	// Deliver session
	return s
}

func getMysqlDB() *sql.DB {
	db, err := sql.Open("mysql", "admin:admin@tcp(localhost:3306)/zencart?charset=utf8")

	if err != nil {
		panic(err)
	}

	return db
}
