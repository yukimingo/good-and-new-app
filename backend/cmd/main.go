package main

import (
	"fmt"
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
		var routes []fiber.Route
		for _, route := range app.GetRoutes(false) {
			fmt.Println(route)
			routes = append(routes, route)
		}
		return c.Status(http.StatusOK).JSON(fiber.Map{"routes": routes})
	})

	app.Get("/user", userController.FindAll)

	app.Get("/user/:email", userController.FindByEmail)

	app.Post("/user", userController.Create)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
