package service

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/repository"
)

type UserService struct {
	repository repository.UserRepositoryInterface
}

type UserServiceInterface interface {
	FindAll() ([]domain.User, error)
}

func NewUserService(ur repository.UserRepositoryInterface) *UserService {
	return &UserService{repository: ur}
}

func (us *UserService) FindAll() ([]domain.User, error) {
	users, err := us.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
