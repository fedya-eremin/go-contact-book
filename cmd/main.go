package main

import (
	"log"
	"s0709-22/internal/service"

	"github.com/gofiber/fiber/v2"
)

func main() {
	api := fiber.New()

	svc := service.New()

	api.Get("/api/v1/contact/:id", svc.GetContact)
	api.Post("/api/v1/contact", svc.PostContact)
	api.Put("/api/v1/contact/:id", svc.PutContact)
	api.Delete("/api/v1/contact/:id", svc.DeleteContact)

	api.Get("/api/v1/group/:id", svc.GetGroup)
	api.Post("/api/v1/group", svc.PostGroup)
	api.Put("/api/v1/group/:id", svc.PutGroup)
	api.Delete("/api/v1/group/:id", svc.DeleteGroup)

	if err := api.Listen(":7080"); err != nil {
		log.Fatalln(err)
	}
}
