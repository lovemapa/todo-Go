package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Todo model
type Todo struct {
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty"`
	User   bson.ObjectId `json:"user" bson:"user"`
	Name   string        `json:"name" bson:"name"`
	Status bool          `json:"status" bson:"status"`
	Date   time.Time     `json:"date" bson:"date"`
}
