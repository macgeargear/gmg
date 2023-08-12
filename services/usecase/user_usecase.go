package usecase

import (
	"github.com/macgeargear/gmg/models"
	repo "github.com/macgeargear/gmg/repositories"
	svc "github.com/macgeargear/gmg/services"
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
func (u *userService) Create(data models.User) error {
	panic("unimplemented")
}

// Delete implements service.UserService.
func (u *userService) Delete(id string) error {
	panic("unimplemented")
}

// GetAll implements service.UserService.
func (u *userService) GetAll() ([]models.User, error) {
	res, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetById implements service.UserService.
func (u *userService) GetById(id string) (*models.User, error) {
	res, err := u.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, err
	}
	return res, nil

}

// Update implements service.UserService.
func (u *userService) Update(data map[string]interface{}, id string) (*models.User, error) {
	panic("unimplemented")
}
