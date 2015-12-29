package main

import (
    // standard library packages
    "net/http"

    // third party packages
    "github.com/julienschmidt/httprouter"

    // project scope packages
    "github.com/shunchaowang/zencart-service/controller"
)

func main() {

    // instantiate a new router
    r := httprouter.New()

    // get a ProductController instance
    pc := controller.NewProductController()

    // get a product
    r.GET("/product/:id", pc.Get)

    // create a product
    r.POST("/product", pc.Create)

    // delete a product
    r.DELETE("/product/:id", pc.Delete)

    // fire up the server
    http.ListenAndServe("localhost:8080", r)
}
