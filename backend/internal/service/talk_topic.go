package service

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/repository"
)

type TalkTopicService struct {
	repository repository.TalkTopicRepositoryInterface
}

type TalkTopicServiceInterface interface {
	FindAll() ([]domain.TalkTopic, error)
	Create(talkTopic domain.TalkTopic) (int64, error)
}

func NewTalkTopicService(ttr repository.TalkTopicRepositoryInterface) *TalkTopicService {
	return &TalkTopicService{repository: ttr}
}

func (tts *TalkTopicService) Create(talkTopic domain.TalkTopic) (int64, error) {
	return tts.repository.Create(talkTopic)
}

func (tts *TalkTopicService) FindAll() ([]domain.TalkTopic, error) {
	return tts.repository.FindAll()
}
