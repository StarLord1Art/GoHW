package db

import (
	"GoHW1/internal/domain/entity"
)

type UserRepository struct {
	storage map[int64]entity.User
}

func (ubr *UserRepository) InitializeUserStorage() {
	ubr.storage = make(map[int64]entity.User)
}

func (ubr *UserRepository) generateId() int64 {
	return int64(len(ubr.storage) + 1)
}

func (ubr *UserRepository) RegisterUser(name string, email string) (entity.User, error) {
	user := entity.User{Name: name, Email: email, Id: ubr.generateId()}
	ubr.storage[user.Id] = user

	return user, nil
}
