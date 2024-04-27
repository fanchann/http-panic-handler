package service

import (
	"http-panic-handler/domain/web"
	"http-panic-handler/repository"
)

type UserService struct {
	userRepo *repository.UsersRepository
}

func NewUserService(repo *repository.UsersRepository) *UserService {
	return &UserService{repo}
}

func (u *UserService) GetUserById(id int) web.WebResponse {
	user, err := u.userRepo.GetUserById(id)
	if err != nil {
		panic(err)
	}
	return web.WebResponse{
		Success: true,
		Data:    user,
	}
}
