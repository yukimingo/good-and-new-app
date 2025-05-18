package usecase

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/service"
)

type TalkTopicUsecase struct {
	service service.TalkTopicServiceInterface
}

type TalkTopicUsecaseInterface interface {
	FindAll() ([]domain.TalkTopic, error)
	Create(talkTopic domain.TalkTopic) (int64, error)
}

func NewTalkTopicUsecase(tts service.TalkTopicServiceInterface) *TalkTopicUsecase {
	return &TalkTopicUsecase{service: tts}
}

func (ttu *TalkTopicUsecase) Create(talkTopic domain.TalkTopic) (int64, error) {
	return ttu.service.Create(talkTopic)
}

func (ttu *TalkTopicUsecase) FindAll() ([]domain.TalkTopic, error) {
	return ttu.service.FindAll()
}
