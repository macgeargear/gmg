package service

import "github.com/macgeargear/gmg/repository"

type studentService struct {
	studentRepo repository.StudentRepo
}

func NewStudentService(studentRepo repository.StudentRepo) studentService {
	return studentService{studentRepo: studentRepo}
}
