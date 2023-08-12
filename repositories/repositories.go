package repositories

import (
	m "github.com/macgeargear/gmg/models"
)

type UserRepo interface {
	GetAll() ([]m.User, error)
	GetBy(filter map[string]interface{}) (*m.User, error)
	Create(user m.User) error
	Update(user *m.User) error
	Delete(userId string) error
}

type LessonRepo interface {
	Find(lessonId string) (*m.Lesson, error)
	Update(Lesson *m.Lesson) error
	FindALl() ([]*m.Lesson, error)
	Delete(lessonId string) error
}

type CourseRepo interface {
	Find(courseId string) (*m.Course, error)
	Update(Course *m.Course) error
	FindALl() ([]*m.Course, error)
	Delete(code string) error
}
