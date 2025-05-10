package domain

import "good-and-new/internal/pkg/timeutil"

type TalkTopic struct {
	ID         uint   `json:"id"`
	CategoryID uint   `json:"category_id"`
	Title      string `json:"title"`
	Content    string `json:"content"`

	timeutil.TimeStamped
}
