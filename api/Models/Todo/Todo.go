package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Todo model
type Todo struct {
	Name   string        `json:"name" bson:"name" binding:"required"`
	ID     bson.ObjectId `json:"id" bson:"_id,omitempty" swaggerignore:"true"`
	User   bson.ObjectId `json:"user" bson:"user"`
	Status bool          `json:"status" bson:"status" swaggerignore:"true"`
	Date   time.Time     `json:"date" bson:"date" swaggerignore:"true"`
}
