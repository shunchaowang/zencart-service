package controller

import (
    // standard library packages
    "encoding/json"
    "fmt"
    "net/http"

    // third party packages
    "github.com/julienschmidt/httprouter"

    // project scope packages
    "github.com/shunchaowang/zencart-service/model"
)

type (
    ProductController struct {}
    ProductTypeController struct {}
)

// ProductController apis
func NewProductController() *ProductController {
    return &ProductController{}
}

func (pc ProductController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // stub an example product
    product := model.Product{
        Id: p.ByName("id"),
        Model: "iPhone",
    }

    // marshal provided interface into JSON structure
    pj, _ := json.Marshal(product)

    // write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(200)
    fmt.Fprintf(w, "%s", pj)
}

func (pc ProductController) Create(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    
    // stub an user to be populated from the body
    product := model.Product{}

    // populate the product data
    json.NewDecoder(r.Body).Decode(&product)

    // add an id
    product.Id = "foo"

    // marsha provided interface into JSON structure
    pj, _ := json.Marshal(product)

    // write content-type, statuscode, payload
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(201)
    fmt.Fprintf(w, "%s", pj)
}

func (pc ProductController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // TODO: only write status for now
    w.WriteHeader(200)
}

// ProductTypeController apis
func (ptc ProductTypeController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    // TODO: 
}
