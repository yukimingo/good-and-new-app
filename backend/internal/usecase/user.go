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
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (int64, error)
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

func (uu *UserUsecase) FindByEmail(email string) (domain.User, error) {
	user, err := uu.service.FindByEmail(email)
	if err != nil {
		return user, err
	}

	return user, err
}

func (uu *UserUsecase) Create(user domain.User) (int64, error) {
	id, err := uu.service.Create(user)
	if err != nil {
		return 0, err
	}

	return id, nil
}
