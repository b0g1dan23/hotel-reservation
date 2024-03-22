package main

import (
	"context"
	"flag"

	"github.com/b0g1dan23/hotel-reservation/api"
	"github.com/b0g1dan23/hotel-reservation/db"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const DBURI = "mongodb://localhost:27017"

var config = fiber.Config{
	ErrorHandler: func(ctx *fiber.Ctx, err error) error {
		return ctx.JSON(map[string]string{"error": err.Error()})
	},
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(DBURI))
	if err != nil {
		panic(err)
	}

	listenAddr := flag.String("listenAddr", "localhost:5000", "The listen address of the API Server")
	flag.Parse()

	app := fiber.New(config)
	apiv1 := app.Group("/api/v1")

	userHandler := api.NewUserHandler(db.NewMongoUserStore(client))
	apiv1.Get("/users", userHandler.HandleGetUsers)
	apiv1.Post("/users", userHandler.HandlePostUser)
	apiv1.Get("/users/:id", userHandler.HandleGetUserByID)
	app.Listen(*listenAddr)
}
