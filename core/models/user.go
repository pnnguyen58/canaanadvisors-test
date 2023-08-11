package models

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int64  `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	RoleId   int64  `json:"role_id,omitempty"`
}
func (u *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	u.Password = string(bytes)
	return nil
}

func (u *User) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

var users = []User{
	{
		Id: 1,
		Username: "user",
		Password: "123456",
		Name: "User",
		RoleId: 1,
	},
	{
		Id: 2,
		Username: "driver",
		Password: "1234567",
		Name: "Driver",
		RoleId: 2,
	},
	{
		Id: 3,
		Username: "owner",
		Password: "12345678",
		Name: "Restaurant owner",
		RoleId: 3,
	},
}

func GetUser(username string) (*User, error) {
	for i := range users {
		if users[i].Username == username {
			return &users[i], nil
		}
	}
	return nil, errors.New("user not found")
}
