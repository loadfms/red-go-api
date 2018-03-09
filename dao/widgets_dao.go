package dao

import (
	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WidgetsDAO struct {
	session *mgo.Session
}

const (
	COLLECTION_WIDGETS = "widgets"
)

func NewWidgetsDAO(session *mgo.Session) *WidgetsDAO {
	return &WidgetsDAO{session}
}

// Find list of users
func (m *WidgetsDAO) FindAll() ([]Widget, error) {
	var widget []Widget
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).Find(bson.M{}).All(&widget)
	return widget, err
}

// Find a user by its id
func (m *WidgetsDAO) FindById(id string) (Widget, error) {
	var widget Widget
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).FindId(bson.ObjectIdHex(id)).One(&widget)
	return widget, err
}

// Insert a user into database
func (m *WidgetsDAO) Insert(widget Widget) error {
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).Insert(&widget)
	return err
}

// Delete an existing user
func (m *WidgetsDAO) Delete(widget Widget) error {
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).Remove(&widget)
	return err
}

// Update an existing user
func (m *WidgetsDAO) Update(widget Widget) error {
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).UpdateId(widget.Id, &widget)
	return err
}
