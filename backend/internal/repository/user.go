package repository

import (
	"database/sql"
	"good-and-new/internal/domain"
)

const (
	selectUsers          = "select * from users"
	selectUserWhereEmail = "select * from users where email = ?"
	insertUsersTable     = "insert into users (name, email, password, created_at, updated_at) values(?, ?, ?, ?, ?)"
)

type UserRepository struct {
	db *sql.DB
}

type UserRepositoryInterface interface {
	FindAll() ([]domain.User, error)
	FindByEmail(email string) (domain.User, error)
	Create(user domain.User) (int64, error)
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (ur *UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	rows, err := ur.db.Query(selectUsers)
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
	if err := ur.db.QueryRow(selectUserWhereEmail, email).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return user, err
	}

	return user, nil
}

func (ur *UserRepository) Create(user domain.User) (int64, error) {
	stmt, err := ur.db.Prepare(insertUsersTable)
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(user.Name, user.Email, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}
