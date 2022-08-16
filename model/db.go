package model

import "gopkg.in/mgo.v2/bson"

type Student struct {
	Id         bson.ObjectId `json:"id "bson:"_id"`
	FirstName  string        `json:"first_name" bson:"first_name"`
	SecondName string        `json:"second_name" bson:"second_name"`
	Faculty    string        `json:"faculty" bson:"faculty"`
	BirthDate  int64         `json:"birth_date" bson:"birth_date"`
}

type Students []Student
