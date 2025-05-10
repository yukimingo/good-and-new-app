package controller

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/pkg/timeutil"
	"good-and-new/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	usecase usecase.UserUsecaseInterface
}

type UserControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	FindByEmail(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
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

func (uc *UserController) FindByEmail(c *fiber.Ctx) error {
	email := c.Params("email")

	user, err := uc.usecase.FindByEmail(email)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
}

func (uc *UserController) Create(c *fiber.Ctx) error {
	var user domain.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	user.SetTimeAt(timeutil.NowJST())

	id, err := uc.usecase.Create(user)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"id": id})
}
