package service_test

import (
	"sync"
	"testing"
	"time"

	m "github.com/macgeargear/gmg/models"
	rh "github.com/macgeargear/gmg/repositories/helper"
	service "github.com/macgeargear/gmg/services"
	"github.com/macgeargear/gmg/services/usecase"
)

var (
	userService service.UserService
)

func UserTestData() []m.User {
	return []m.User{
		{
			UserId:   "test1",
			Name:     "cookie",
			Email:    "cookie@gmail.com",
			Password: "1234",
		},
		{
			UserId:   "test2",
			Name:     "cookie2",
			Email:    "cookie2@gmail.com",
			Password: "1234",
		},
		{
			UserId:   "test3",
			Name:     "cookie3",
			Email:    "cookie3@gmail.com",
			Password: "1234",
		},
		{
			UserId:   "test4",
			Name:     "cookie4",
			Email:    "cookie4@gmail.com",
			Password: "1234",
		},
	}
}

func init() {
	userRepo := rh.ChooseRepo()
	userService = usecase.NewUserService(userRepo)
}

func TestUserService(t *testing.T) {
	t.Run("Insert User", InsertUser)
}

func InsertUser(t *testing.T) {
	testdata := UserTestData()
	wg := sync.WaitGroup{}

	// Clean test data if any
	for _, data := range testdata {
		wg.Add(1)
		go func(_data m.User) {
			userService.Delete(_data.UserId)
			wg.Done()
		}(data)
	}

	wg.Wait()
	time.Sleep(time.Second * 1)

	t.Run("Case 1: Save data", func(t *testing.T) {
		for _, data := range testdata {
			wg.Add(1)
			go func(_data m.User) {
				err := userService.Create(_data)
				if err != nil {
					t.Errorf("[ERROR] - Failed to save dta %s", err.Error())
				}
				wg.Done()
			}(data)
		}
		wg.Wait()
		time.Sleep(time.Second * 1)

		for _, data := range testdata {
			res, err := userService.GetById(data.UserId)
			if err != nil || res.UserId == "" {
				t.Errorf("[ERROR] - failed to get data")
			}
		}
	})

	t.Run("Case 2: Negative Test", func(t *testing.T) {
		t.Run("Case 2.1: Duplicate email", func(t *testing.T) {
			_data := testdata[0]
			_data.Email = "cookie4@gmail.com"
			err := userService.Create(_data)
			if err == nil {
				t.Error("[ERROR] - duplicate validation email is not working")
			}
		})
		t.Run("Case 2.2: Duplicate id", func(t *testing.T) {
			_data := testdata[0]
			_data.UserId = "test4"
			err := userService.Create(_data)
			if err == nil {
				t.Error("[ERROR] - duplicate validation userId is not working")
			}
		})
	})
}
