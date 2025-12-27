package repository

import "GoHW1/internal/domain/entity"

type UserRepository interface {
	RegisterUser(name string, email string) (entity.User, error)
	InitializeUserStorage()
}
