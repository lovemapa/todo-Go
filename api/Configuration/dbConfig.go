package configuration

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

// ConnectDb is database configuration
func ConnectDb(Db string) (mongoSession *mgo.Session) {
	mongoDBDialInfo := &mgo.DialInfo{

		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Database: Db,
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession: %s\n", err)
	}
	mongoSession.SetMode(mgo.Monotonic, true)

	return mongoSession

}
