package repository

import (
	"database/sql"
	"good-and-new/internal/domain"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := ur.db.Query("select * from users")
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var user domain.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
			return users, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (ur *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	if err := ur.db.QueryRow("select * from users where email = ?", email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}

	return user, nil
}
