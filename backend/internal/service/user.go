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
	Create(user domain.User) error
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

func (us *UserService) Create(user domain.User) error {
	if err := us.repository.Create(user); err != nil {
		return err
	}

	return nil
}
