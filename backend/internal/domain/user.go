package domain

import "good-and-new/internal/pkg/timeutil"

type User struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`

	timeutil.TimeStamped
}
