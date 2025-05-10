package controller

import (
	"good-and-new/internal/domain"
	"good-and-new/internal/pkg/timeutil"
	"good-and-new/internal/usecase"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type CategoryController struct {
	usecase usecase.CategoryUsecaseInterface
}

type CategoryControllerInterface interface {
	FindAll(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
}

func NewCategoryController(cu usecase.CategoryUsecaseInterface) *CategoryController {
	return &CategoryController{usecase: cu}
}

func (cc *CategoryController) FindAll(c *fiber.Ctx) error {
	categories, err := cc.usecase.FindAll()
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"categories": categories})
}

func (cc *CategoryController) Create(c *fiber.Ctx) error {
	var category domain.Category
	if err := c.BodyParser(&category); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	category.SetTimeAt(timeutil.NowJST())

	id, err := cc.usecase.Create(category)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{"id": id})
}
