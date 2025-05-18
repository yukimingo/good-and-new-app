package main

import (
	"good-and-new/internal/controller"
	"good-and-new/internal/db"
	"good-and-new/internal/repository"
	"good-and-new/internal/service"
	"good-and-new/internal/usecase"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService)
	userController := controller.NewUserController(userUsecase)

	categoryRepository := repository.NewCategoryRepository(db)
	categoryService := service.NewCategoryService(categoryRepository)
	categoryUsecase := usecase.NewCategoryUsecase(categoryService)
	categoryController := controller.NewCategoryController(categoryUsecase)

	talkTopicRepository := repository.NewTalkTopicRepository(db)
	talkTopicService := service.NewTalkTopicService(talkTopicRepository)
	talkTopicUsecase := usecase.NewTalkTopicUsecase(talkTopicService)
	talkTopicController := controller.NewTalkTopicController(talkTopicUsecase)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	app.Get("/user", userController.FindAll)

	app.Get("/user/:email", userController.FindByEmail)

	app.Post("/user", userController.Create)

	app.Get("/category", categoryController.FindAll)

	app.Post("/category", categoryController.Create)

	app.Get("/topic", talkTopicController.FindAll)

	app.Post("/topic", talkTopicController.Create)

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}
