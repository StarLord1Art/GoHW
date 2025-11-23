package handler

import (
	"GoHW1/internal/application/dto/request"
	"GoHW1/internal/application/usecase"
	"encoding/json"
	"log"
	"net/http"
)

type UserHandler struct {
	userService usecase.UserService
}

func (uh *UserHandler) InitializeUserStorage() {
	uh.userService.InitializeUserStorage()
}

func (uh *UserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var body request.UserRequestDto
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		http.Error(w, "incorrect data format", http.StatusBadRequest)
		return
	}

	user, err := uh.userService.RegisterUser(body.Name, body.Email)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Println(err)
		return
	}
}
