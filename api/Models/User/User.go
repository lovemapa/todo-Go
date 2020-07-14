package models

import "gopkg.in/mgo.v2/bson"

// User model
type User struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName      string        `json:"firstName" bson:"firstName"`
	LastName       string        `json:"lastName" bson:"lastName"`
	Email          string        `json:"email" bson:"email"`
	Token          string        `json:"token" bson:"token"`
	Mobile         string        `json:"mobile" bson:"mobile"`
	Password       string        `json:"password" bson:"password"`
	ProfilePhoto   string        `json:"ProfilePhoto" bson:"ProfilePhoto" binding:"required"`
	Date           int64         `json:"date" bson:"date"`
	TokenExpiresAt int64         `json:"tokenExpiresAt" bson:"tokenExpiresAt"`
}

//UserLogin model
type UserLogin struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	Token    string `json:"token" bson:"token"`
}
