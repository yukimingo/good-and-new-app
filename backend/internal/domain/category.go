package domain

import (
	"good-and-new/internal/pkg/timeutil"
)

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`

	timeutil.TimeStamped
}
