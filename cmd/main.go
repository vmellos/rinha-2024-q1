package main

import (
	"context"
	"rinha_2024_q1/internal/adapter/input/controller"
	"rinha_2024_q1/internal/adapter/input/controller/routes"
	"rinha_2024_q1/internal/adapter/output/repository"
	"rinha_2024_q1/internal/application/service"
	mongodb "rinha_2024_q1/internal/configuration/database/mongo"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	database := mongodb.NewMongoDBConnection(context.Background())

	transactionController := initDependencies(database)
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	routes.InitRoutes(app, transactionController)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func initDependencies(database *mongo.Database) controller.TransactionControllerInterface {
	transactionRepo := repository.NewTransactionRepository(database)
	transactionService := service.NewTransactionDomainService(transactionRepo)
	return controller.NewTransactionControllerInterface(transactionService)
}
