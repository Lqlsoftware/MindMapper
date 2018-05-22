package orm

import (
	"github.com/Lqlsoftware/mindmapper/src/config"
	"gopkg.in/mgo.v2"
)

type Database struct {
	session	*mgo.Session
	Db		*mgo.Database
}

func (orm *Database)GetSession() *mgo.Session {
	return orm.session.Copy()
}

var ORM *Database

func InitDB () {
	// Connect mongodb
	session, err := mgo.Dial(config.DB_URL)
	if err != nil {
		panic(err)
	}

	// Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	// Get DB pointer
	Db := session.DB(config.DB_NAME)

	// Set globe variable
	ORM = &Database{
		session: 	session,
		Db:			Db,
	}
}