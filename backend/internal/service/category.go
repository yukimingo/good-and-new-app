package service

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/repository"
)

type CategoryService struct {
	repository repository.CategoryRepositoryInterface
}

type CategoryServiceInterface interface {
	FindAll() ([]domain.Category, error)
	Create(category domain.Category) (int64, error)
}

func NewCategoryService(cr repository.CategoryRepositoryInterface) *CategoryService {
	return &CategoryService{repository: cr}
}

func (cs *CategoryService) FindAll() ([]domain.Category, error) {
	categories, err := cs.repository.FindAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}

func (cs *CategoryService) Create(category domain.Category) (int64, error) {
	return cs.repository.Create(category)
}
