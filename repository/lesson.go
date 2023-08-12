package repository

type Lesson struct {
	LessonId string `json:"lessonId" bson:"lessonId"`
	Name     string `json:"name" bson:"name"`
	Url      string `json:"url" bson:"url"`
}

type LessonRepo interface {
	Find(lessonId string) (*Lesson, error)
	Update(Lesson *Lesson) error
	FindALl() ([]*Lesson, error)
	Delete(lessonId string) error
}
