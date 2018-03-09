package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id" json:"id"`
	Name     string        `bson:"name" json:"name"`
	Gravatar string        `bson:"gravatar" json:"gravatar"`
}
