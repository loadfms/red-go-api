package dao

import (
	"log"

	. "../config"

	mgo "gopkg.in/mgo.v2"
)

var session *mgo.Session
var config = Config{}

func GetSession() *mgo.Session {
	config.Read()
	if session == nil {
		var e error
		session, e = mgo.Dial(config.Server)
		if e != nil {
			log.Fatalln("----- Failed to connect -----")
			log.Fatalln(e)
		} else {
			log.Println("----- Connected to mongo -----")
		}
	}

	return session
}
