package controller

import (
	"rinha_2024_q1/internal/application/port/input"

	"github.com/gofiber/fiber/v2"
)

func NewTransactionControllerInterface(serviceInterface input.TransactionDomainService) TransactionControllerInterface {
	return &transactionControllerInterface{service: serviceInterface}
}

type TransactionControllerInterface interface {
	TransactionHandler(c *fiber.Ctx) error
	ExtractHandler(c *fiber.Ctx) error
}

type transactionControllerInterface struct {
	service input.TransactionDomainService
}

func (tc *transactionControllerInterface) TransactionHandler(c *fiber.Ctx) error {
	return tc.service.TransactionHandler(c)
}

func (tc *transactionControllerInterface) ExtractHandler(c *fiber.Ctx) error {
	return tc.service.ExtractHandler(c)
}
