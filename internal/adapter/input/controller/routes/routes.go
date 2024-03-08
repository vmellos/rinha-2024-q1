package routes

import (
	"rinha_2024_q1/internal/adapter/input/controller"

	"github.com/gofiber/fiber/v2"
)

func InitRoutes(app *fiber.App, transactionController controller.TransactionControllerInterface) {

	app.Post("/clientes/:id/transacoes", transactionController.TransactionHandler)
	app.Get("/clientes/:id/extrato", transactionController.ExtractHandler)
}
