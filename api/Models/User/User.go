package models

import "gopkg.in/mgo.v2/bson"

// User model
type User struct {
	ID             bson.ObjectId `json:"id" bson:"_id,omitempty"`
	FirstName      string        `json:"firstName" bson:"firstName" binding:"required"`
	LastName       string        `json:"lastName" bson:"lastName" binding:"required"`
	Email          string        `json:"email" bson:"email" binding:"required"`
	Token          string        `json:"token" bson:"token"`
	Mobile         string        `json:"mobile" bson:"mobile" binding:"required"`
	Password       string        `json:"password" bson:"password" binding:"required"`
	ProfilePhoto   string        `json:"ProfilePhoto" bson:"ProfilePhoto"`
	Date           int64         `json:"date" bson:"date"`
	TokenExpiresAt int64         `json:"tokenExpiresAt" bson:"tokenExpiresAt"`
}

//UserLogin model
type UserLogin struct {
	Email    string `json:"email" bson:"email" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
	Token    string `json:"token" bson:"token"`
}
