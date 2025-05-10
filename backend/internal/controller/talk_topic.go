package controller

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/pkg/timeutil"
	"good-and-new/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type TalkTopicController struct {
	usecase usecase.TalkTopicUsecaseInterface
}

type TalkTopicControllerInterface interface {
	Create(c *fiber.Ctx) error
}

func NewTalkTopicController(ttu usecase.TalkTopicUsecaseInterface) *TalkTopicController {
	return &TalkTopicController{usecase: ttu}
}

func (ttc *TalkTopicController) Create(c *fiber.Ctx) error {
	var talkTopic domain.TalkTopic
	if err := c.BodyParser(&talkTopic); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	talkTopic.SetTimeAt(timeutil.NowJST())

	id, err := ttc.usecase.Create(talkTopic)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"id": id})
}
