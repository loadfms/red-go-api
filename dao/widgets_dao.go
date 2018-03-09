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

// Find list of widgets
func (m *WidgetsDAO) FindAll() ([]Widget, error) {
	var widget []Widget
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).Find(bson.M{}).All(&widget)
	return widget, err
}

// Find a widget by its id
func (m *WidgetsDAO) FindById(id string) (Widget, error) {
	var widget Widget
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).FindId(bson.ObjectIdHex(id)).One(&widget)
	return widget, err
}

// Insert a widget into database
func (m *WidgetsDAO) Insert(widget Widget) error {
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).Insert(&widget)
	return err
}

// Delete an existing widget
func (m *WidgetsDAO) Delete(id string) error {
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).RemoveId(bson.ObjectIdHex(id))
	return err
}

// Update an existing widget
func (m *WidgetsDAO) Update(id string, widget Widget) error {
	widget.Id = bson.ObjectIdHex(id)
	err := session.DB(config.Database).C(COLLECTION_WIDGETS).UpdateId(bson.ObjectIdHex(id), &widget)
	return err
}
