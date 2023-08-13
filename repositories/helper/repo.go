package helper

import (
	"log"
	"os"
	"strconv"

	repo "github.com/macgeargear/gmg/repositories"
	mg "github.com/macgeargear/gmg/repositories/mongodb"
)

func ChooseRepo() repo.UserRepo {
	url := os.Getenv("url")
	db := os.Getenv("db")
	timeout, _ := strconv.Atoi(os.Getenv("timeout"))

	if url == "" {
		url = "mongodb://localhost:27017/"
	}
	if db == "" {
		db = "local"
	}
	if timeout == 0 {
		timeout = 10
	}

	repo, err := mg.NewUserRepository(url, db, timeout)
	if err != nil {
		log.Fatal(err)
	}
	return repo
}
