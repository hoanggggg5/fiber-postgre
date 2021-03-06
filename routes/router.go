package routes

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/hoanggggg5/fiber-postgre/controllers/identity"
	"github.com/hoanggggg5/fiber-postgre/controllers/public"
	"github.com/hoanggggg5/fiber-postgre/controllers/resource"
)

func InitRouter() {
	app := fiber.New()

	app.Get("/public/timestamp", public.GetTimestamp)
	app.Get("/resource/users/me", resource.GetUser)
	app.Post("/resource/users/create", resource.CreateUser)
	app.Post("/resource/users/login", identity.Login)

	log.Fatal(app.Listen(":3001"))
}
