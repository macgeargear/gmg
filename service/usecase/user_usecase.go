package usecase

import (
	"github.com/macgeargear/gmg/models"
	repo "github.com/macgeargear/gmg/repositories"
	svc "github.com/macgeargear/gmg/service"
)

type userService struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) svc.UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Create implements service.UserService.
func (*userService) Create(data models.User) error {
	panic("unimplemented")
}

// Delete implements service.UserService.
func (*userService) Delete(id string) error {
	panic("unimplemented")
}

// GetAll implements service.UserService.
func (*userService) GetAll() ([]models.User, error) {
	panic("unimplemented")
}

// GetById implements service.UserService.
func (*userService) GetById(id string) (*models.User, error) {
	panic("unimplemented")
}

// Update implements service.UserService.
func (*userService) Update(data map[string]interface{}, id string) (*models.User, error) {
	panic("unimplemented")
}
