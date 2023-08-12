package models

type User struct {
	UserId   string `json:"userId" bson:"userId"`
	Name     string `json:"name" bson:"name"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func NewUser() *User {
	m := new(User)
	return m
}

func (m *User) TableName() string {
	return "users"
}

func NewUserDefaultData() *User {
	user := NewUser()
	user.Name = "cookie"
	user.UserId = "u0001"
	user.Email = "foo@bar.com"
	user.Password = "password"
	return user
}

func (user *User) GetMapFormat() map[string]interface{}
