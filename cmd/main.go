package main

import (
	"context"
	"rinha_2024_q1/internal/adapter/input/controller"
	"rinha_2024_q1/internal/adapter/input/controller/routes"
	"rinha_2024_q1/internal/adapter/output/repository"
	"rinha_2024_q1/internal/application/port/output"
	"rinha_2024_q1/internal/application/service"
	mongodb "rinha_2024_q1/internal/configuration/database/mongo"
	"rinha_2024_q1/internal/configuration/database/postgres"

	"github.com/gofiber/fiber/v2"
)

var (
	repo string
)

func main() {

	transactionController := initDependencies()
	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	routes.InitRoutes(app, transactionController)

	if err := app.Listen(":8080"); err != nil {
		panic(err)
	}
}

func initDependencies() controller.TransactionControllerInterface {
	var transactionRepo output.TransactionPort
	if repo == "mongo" {
		transactionRepo = StartMongoRepo()
	} else {
		transactionRepo = StartPostgresRepo()
	}
	transactionService := service.NewTransactionDomainService(transactionRepo)
	return controller.NewTransactionControllerInterface(transactionService)
}

func StartMongoRepo() output.TransactionPort {
	database := mongodb.NewMongoDBConnection(context.Background())
	transactionRepo := repository.NewTransactionMongoRepository(database)
	return transactionRepo
}

func StartPostgresRepo() output.TransactionPort {
	database, _ := postgres.OpenConn()
	transactionRepo := repository.NewTransactionPostgresRepository(database)
	return transactionRepo
}
