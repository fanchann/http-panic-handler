package handler

import (
	"encoding/json"
	"http-panic-handler/service"
	"net/http"
	"strconv"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service}
}

func (u *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}

	user := u.userService.GetUserById(idInt)
	byteJson, err := json.MarshalIndent(&user, " ", " ")
	if err != nil {
		panic(err)
	}

	w.Write(byteJson)
}
