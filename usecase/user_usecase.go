package usecase

import (
	"go-rest-api/model"
	"go-rest-api/repository"
)

type IUserUsecase interface {
	SignUp(user model.User) (model.UserResponse, error)
	Login(user model.User) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
}