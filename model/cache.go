package model

type CacheStudent struct {
	Id         string `json:"id" bson:"_id"`
	FirstName  string `json:"first_name" bson:"first_name"`
	SecondName string `json:"second_name" bson:"second_name"`
	Faculty    string `json:"faculty" bson:"faculty"`
	BirthDate  string `json:"birth_date" bson:"birth_date"`
}

type CacheStudents []CacheStudent
