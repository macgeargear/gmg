package repository

type Course struct {
	CourseId string  `json:"courseId" bson:"courseId"`
	Name     string  `json:"name" bson:"name"`
	Price    float32 `json:"price" bson:"price"`
}

type CourseRepo interface {
	Find(courseId string) (*Course, error)
	Update(Course *Course) error
	FindALl() ([]*Course, error)
	Delete(code string) error
}
