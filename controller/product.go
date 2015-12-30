package controller

import (
	// standard library packages
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	// third party packages
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	// project scope packages
	"github.com/shunchaowang/zencart-service/model"
)

type (
	ProductController struct {
		session *mgo.Session
	}
	ProductTypeController struct{}
)

// ProductController apis
func NewProductController(s *mgo.Session) *ProductController {
	return &ProductController{s}
}

func (pc ProductController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab id
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab object id
	oid := bson.ObjectIdHex(id)

	// Stub Product
	product := model.Product{}

	// Fetch Product
	if err := pc.session.DB("zencart").C("products").FindId(oid).One(&product); err != nil {
		w.WriteHeader(404)
		return
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
	product.Id = bson.NewObjectId()

	// Write the product to mongo
	pc.session.DB("zencart").C("products").Insert(product)

	// marsha provided interface into JSON structure
	pj, _ := json.Marshal(product)

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	fmt.Fprintf(w, "%s", pj)
}

func (pc ProductController) Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// Grab param
	id := p.ByName("id")

	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	// Grab id
	oid := bson.ObjectIdHex(id)

	// Delete Product
	if err := pc.session.DB("zencart").C("products").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	// Write status
	w.WriteHeader(200)
}

// ProductTypeController apis
func (ptc ProductTypeController) Get(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO:
}
