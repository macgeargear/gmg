package models

type Lesson struct {
	LessonId string `json:"lessonId" bson:"lessonId"`
	Name     string `json:"name" bson:"name"`
	Url      string `json:"url" bson:"url"`
}
