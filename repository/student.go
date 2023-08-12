package repository

type Student struct {
	StudentId string `json:"studentId" bson:"studentId"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
	Password  string `json:"password" bson:"password"`
}

type StudentRepo interface {
	Find(studentId string) (*Student, error)
	Update(student *Student) error
	FindAll() ([]Student, error)
	Delete(studentId string) error
}
