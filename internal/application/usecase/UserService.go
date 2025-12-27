package usecase

import (
	"GoHW1/internal/application/dto/response"
	"GoHW1/internal/infrastructure/db"
	"errors"
)

type UserService struct {
	userRepository db.UserRepository
}

func (u *UserService) InitializeUserStorage() {
	u.userRepository.InitializeUserStorage()
}

func (u *UserService) RegisterUser(name string, email string) (response.UserResponseDto, error) {
	if name == "" {
		return response.UserResponseDto{}, errors.New("name must not be empty")
	} else if email == "" {
		return response.UserResponseDto{}, errors.New("email must not be empty")
	}

	registeredUser, err := u.userRepository.RegisterUser(name, email)
	if err != nil {
		return response.UserResponseDto{}, err
	}

	return response.UserResponseDto{Id: registeredUser.Id, Name: registeredUser.Name, Email: registeredUser.Email}, nil
}
