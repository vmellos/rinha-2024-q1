package service

import (
	"rinha_2024_q1/internal/application/domain"

	"github.com/gofiber/fiber/v2"
)

func (td *transactionDomainService) ExtractHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	cliente := domain.ClienteDomain{
		Id: id,
	}
	clientDomainResult := td.repository.GetClient(cliente)

	if clientDomainResult == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	transactionDomainResult := td.repository.Find(cliente)

	return c.Status(fiber.StatusOK).JSON(transactionDomainResult)
}
