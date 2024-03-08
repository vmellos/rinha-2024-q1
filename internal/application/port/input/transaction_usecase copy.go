package input

import "github.com/gofiber/fiber/v2"

type TransactionDomainService interface {
	TransactionHandler(c *fiber.Ctx) error
	ExtractHandler(c *fiber.Ctx) error
}
