package usecase

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/service"
)

type UserUsecase struct {
	service service.UserServiceInterface
}

type UserUsecaseInterface interface {
	FindAll() ([]domain.User, error)
}

func NewUserUsecase(us service.UserServiceInterface) *UserUsecase {
	return &UserUsecase{service: us}
}

func (uu *UserUsecase) FindAll() ([]domain.User, error) {
	users, err := uu.service.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
