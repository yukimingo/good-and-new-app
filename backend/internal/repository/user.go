package repository

import (
	"good-and-new/internal/domain"
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) error
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User

	result := ur.db.Find(&users)
	log.Println("SQL:", result.Statement.SQL.String())

	return users, result.Error
}

func (ur *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User

	if row := ur.db.Where("email = ?", email).First(&user); row.Error != nil {
		return user, row.Error
	}

	return user, nil
}

func (ur *UserRepository) Create(user domain.User) error {
	if err := ur.db.Create(&user).Error; err != nil {
		return err
	}

	return nil
}
