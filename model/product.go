package model

import "gopkg.in/mgo.v2/bson"

type (
	Product struct {
		Id    bson.ObjectId `json:"id" bson:"_id"`
		Model string        `json:"model" bson:"model"`
	}

	ProductType struct {
		Id   bson.ObjectId `json:"id" bson:"_id"`
		Name string        `json:"name" bson:"name"`
	}
)
