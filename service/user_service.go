package service

import (
	m "github.com/macgeargear/gmg/models"
)

type UserService interface {
	GetAll() ([]m.User, error)
	GetById(id string) (*m.User, error)
	Create(data m.User) error
	Update(data map[string]interface{}, id string) (*m.User, error)
	Delete(id string) error
}
