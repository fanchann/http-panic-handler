package main

import (
	"http-panic-handler/handler"
	"http-panic-handler/middleware"
	"http-panic-handler/repository"
	"http-panic-handler/service"
	"net/http"
)

func main() {
	userRepo := repository.NewUsersRepository()
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	mux := http.ServeMux{}
	mux.HandleFunc("/{id}", userHandler.GetUserById)

	http.ListenAndServe(":9000", middleware.PanicRecovery(&mux))
}
