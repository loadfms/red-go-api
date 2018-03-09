package models

import "gopkg.in/mgo.v2/bson"

type Widget struct {
	Id        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Color     string        `bson:"color" json:"color"`
	Price     string        `bson:"price" json:"price"`
	Inventory int           `bson:"inventory" json:"inventory"`
	Melts     bool          `bson:"melts" json:"melts"`
}
