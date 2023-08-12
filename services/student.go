package service

type StudentRequest struct {
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

type StudentResponse struct {
	StudentId string `json:"studentId" bson:"studentId"`
	Name      string `json:"name" bson:"name"`
	Email     string `json:"email" bson:"email"`
}

type StudentService interface {
	GetStudents() ([]StudentResponse, error)
	GetStudent(string) (*StudentResponse, error)
	Delete(string) (*StudentResponse, error)
}
