package main

import (
	"log"
	"rinha_2024_q1/internal"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Post("/clientes/:id/transacoes", internal.TransacaoHandler)
	app.Get("/clientes/:id/extrato", internal.ExtratoHandler)
	log.Fatal(app.Listen(":8080"))
}
