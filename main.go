package main

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
	"strconv"
)

func middleware(c fiber.Ctx) error {
	return c.Next()
}

func getUsers(c fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	user := User{
		Id:   id,
		Name: "John",
	}
	return c.JSON(user)
}

func main() {
	app := fiber.New()

	app.Use(logger.New())

	api := app.Group("/api", middleware)
	v1 := api.Group("/v1", middleware)
	v1.Get("/users/:id", getUsers)

	log.Fatal(app.Listen(":8080"))
}

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
