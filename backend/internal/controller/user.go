package controller

import (
	"good-and-new/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase usecase.UserUsecaseInterface
}

type UserControllerInterface interface {
	FindAll(c *fiber.Ctx) error
}

func NewUserController(uu usecase.UserUsecaseInterface) *UserController {
	return &UserController{usecase: uu}
}

func (uc *UserController) FindAll(c *fiber.Ctx) error {
	users, err := uc.usecase.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"users": users})
}
