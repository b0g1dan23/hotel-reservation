package main

import (
	"flag"

	"github.com/b0g1dan23/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenAddr", "localhost:5000", "The listen address of the API Server")
	flag.Parse()

	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUserByID)
	app.Listen(*listenAddr)
}
