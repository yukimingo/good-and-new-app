package usecase

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/service"
)

type CategoryUsecase struct {
	service service.CategoryServiceInterface
}

type CategoryUsecaseInterface interface {
	FindAll() ([]domain.Category, error)
	Create(category domain.Category) (int64, error)
}

func NewCategoryUsecase(cs service.CategoryServiceInterface) *CategoryUsecase {
	return &CategoryUsecase{service: cs}
}

func (cu *CategoryUsecase) FindAll() ([]domain.Category, error) {
	categories, err := cu.service.FindAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (cu *CategoryUsecase) Create(category domain.Category) (int64, error) {
	return cu.service.Create(category)
}
