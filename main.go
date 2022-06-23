package main

import (
	"github.com/deukyunlee/protocol-camp/user"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Project!!")
}

func Routers(app *fiber.App) {
	app.Get("/users", user.GetUsers)
	// app.Get("/", user.GetUsers2)
	app.Get("/user/:id", user.GetUser)
	// app.Get("/usertest/:id", user.GetUser2)
	app.Post("/user", user.SaveUser)
	app.Delete("/user/:id", user.DeleteUser)
	app.Put("/user/:id", user.UpdateUser)
}

func main() {
	// result := test.Test_plus(3,5)
	// fmt.Println(result)
	user.InitialMigration()
	app := fiber.New()
	app.Get("/", hello)
	Routers(app)
	app.Listen(":80")
}