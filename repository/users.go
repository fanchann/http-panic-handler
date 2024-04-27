package repository

import (
	"errors"
	"http-panic-handler/domain/models"
)

var ErrUserNotFound = errors.New("user not found")

var fakeData = []models.Users{
	{ID: 1, Username: "Farda"},
	{ID: 2, Username: "Fanny"},
	{ID: 3, Username: "Udin"},
}

type UsersRepository struct{}

func NewUsersRepository() *UsersRepository {
	return &UsersRepository{}
}

func (u *UsersRepository) GetUserById(id int) (*models.Users, error) {
	var response models.Users
	for _, user := range fakeData {
		if id == user.ID {
			response = user
			return &response, nil
		}
	}
	return nil, ErrUserNotFound
}
