package dao

import (
	. "../models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UsersDAO struct {
	session *mgo.Session
}

const (
	COLLECTION_USERS = "users"
)

func NewUsersDAO(session *mgo.Session) *UsersDAO {
	return &UsersDAO{session}
}

// Find list of users
func (m *UsersDAO) FindAll() ([]User, error) {
	var users []User
	err := session.DB(config.Database).C(COLLECTION_USERS).Find(bson.M{}).All(&users)
	return users, err
}

// Find a user by its id
func (m *UsersDAO) FindById(id string) (User, error) {
	var user User
	err := session.DB(config.Database).C(COLLECTION_USERS).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}

// Insert a user into database
func (m *UsersDAO) Insert(user User) error {
	err := session.DB(config.Database).C(COLLECTION_USERS).Insert(&user)
	return err
}

// Delete an existing user
func (m *UsersDAO) Delete(user User) error {
	err := session.DB(config.Database).C(COLLECTION_USERS).Remove(&user)
	return err
}

// Update an existing user
func (m *UsersDAO) Update(user User) error {
	err := session.DB(config.Database).C(COLLECTION_USERS).UpdateId(user.Id, &user)
	return err
}
