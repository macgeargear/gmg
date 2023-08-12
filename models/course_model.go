package models

type Course struct {
	CourseId string  `json:"courseId" bson:"courseId"`
	Name     string  `json:"name" bson:"name"`
	Price    float32 `json:"price" bson:"price"`
}
