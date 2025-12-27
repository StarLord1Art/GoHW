package service

import "GoHW1/internal/application/dto/response"

type UserService interface {
	RegisterUser(name string, email string) (response.UserResponseDto, error)
	InitializeUserStorage()
}
