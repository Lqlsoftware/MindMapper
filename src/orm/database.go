package orm

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"gopkg.in/mgo.v2"
)

func GetDatabase() *mgo.Database {
	return ORM.Copy().DB(config.DB_NAME)
}

var ORM *mgo.Session

func InitDB () {
	// Connect mongodb
	session, err := mgo.Dial(config.DB_URL)
	if err != nil {
		panic(err)
	}

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Set globe variable
	ORM = session
}