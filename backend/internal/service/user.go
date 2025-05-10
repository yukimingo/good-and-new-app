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
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (int64, error)
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

func (us *UserService) FindByEmail(email string) (domain.User, error) {
	user, err := us.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (us *UserService) Create(user domain.User) (int64, error) {
	id, err := us.repository.Create(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
