package usecase

import (
	"errors"

	"github.com/macgeargear/gmg/helper"
	m "github.com/macgeargear/gmg/models"
	repo "github.com/macgeargear/gmg/repositories"
	svc "github.com/macgeargear/gmg/services"
	"github.com/teris-io/shortid"
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
func (u *userService) Create(data m.User) error {
	// TODO: validate
	if data.UserId == "" {
		data.UserId = shortid.MustGenerate()
	}
	if isFound, _ := u.GetById(data.UserId); isFound != nil {
		return errors.New(helper.ErrUserNameDuplicate.Error())
	}
	// * Encrypt password
	return u.userRepo.Create(data)
}

// Delete implements service.UserService.
func (u *userService) Delete(id string) error {
	if id == "" {
		return errors.New("user invalid")
	}
	if err := u.userRepo.Delete(id); err != nil {
		return err
	}
	return nil
}

// GetAll implements service.UserService.
func (u *userService) GetAll() ([]m.User, error) {
	res, err := u.userRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetById implements service.UserService.
func (u *userService) GetById(id string) (*m.User, error) {
	res, err := u.userRepo.GetBy(map[string]interface{}{"ID": id})
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Update implements service.UserService.
func (u *userService) Update(data map[string]interface{}, id string) (*m.User, error) {
	panic("unimplemented")
}
