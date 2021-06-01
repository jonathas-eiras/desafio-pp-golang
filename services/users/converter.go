package users

import (
	"github.com/brianvoe/gofakeit/v6"
	_ "github.com/brianvoe/gofakeit/v6"
)

func NewUser() UserRequest {
	return UserRequest{
		Name:   gofakeit.Name(),
		Email:  gofakeit.Email(),
		Gender: "Male",
		Status: "Active",
	}

}

func UpdateUser() UserRequestUpdate {
	return UserRequestUpdate{
		Name: gofakeit.Name(),
	}
}
