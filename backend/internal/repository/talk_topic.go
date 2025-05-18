package repository

import (
	"database/sql"
	"good-and-new/internal/domain"
)

type TalkTopicRepository struct {
	db *sql.DB
}

type TalkTopicRepositoryInterface interface {
	FindAll() ([]domain.TalkTopic, error)
	Create(talkTopic domain.TalkTopic) (int64, error)
}

func NewTalkTopicRepository(db *sql.DB) *TalkTopicRepository {
	return &TalkTopicRepository{db: db}
}

func (ttr *TalkTopicRepository) Create(talkTopic domain.TalkTopic) (int64, error) {
	stmt, err := ttr.db.Prepare("insert into talk_topics (category_id, title, content, created_at, updated_at) values(?, ?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(talkTopic.CategoryID, talkTopic.Title, talkTopic.Content, talkTopic.CreatedAt, talkTopic.UpdatedAt)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ttr *TalkTopicRepository) FindAll() ([]domain.TalkTopic, error) {
	var topics []domain.TalkTopic
	rows, err := ttr.db.Query("select * from talk_topics")
	if err != nil {
		return topics, err
	}

	for rows.Next() {
		var topic domain.TalkTopic
		if err := rows.Scan(&topic.ID, &topic.CategoryID, &topic.Title, &topic.Content, &topic.CreatedAt, &topic.UpdatedAt); err != nil {
			return topics, err
		}

		topics = append(topics, topic)
	}

	return topics, nil
}
