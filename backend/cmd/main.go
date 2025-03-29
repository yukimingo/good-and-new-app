package main

import (
	"good-and-new/internal/controller"
	"good-and-new/internal/db"
	"good-and-new/internal/repository"
	"good-and-new/internal/service"
	"good-and-new/internal/usecase"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userController := controller.NewUserController(userUsecase)

	app := fiber.New()

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(fiber.Map{"msg": "test"})
	})

	app.Get("/user", userController.FindAll)

	app.Get("/user/:email", func(c *fiber.Ctx) error {
		email := c.Params("email")
		user, err := userRepository.FindByEmail(email)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{"user": user})
	})

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
