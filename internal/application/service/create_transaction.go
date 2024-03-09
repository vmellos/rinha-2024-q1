package service

import (
	"rinha_2024_q1/internal/adapter/input/model/request"
	"rinha_2024_q1/internal/application/domain"
	"rinha_2024_q1/internal/application/port/input"
	"rinha_2024_q1/internal/application/port/output"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewTransactionDomainService(transactionRepository output.TransactionPort) input.TransactionDomainService {
	return &transactionDomainService{transactionRepository}
}

type transactionDomainService struct {
	repository output.TransactionPort
}

func (td *transactionDomainService) TransactionHandler(c *fiber.Ctx) error {
	id := c.Params("id")
	var t request.Transacao
	err := c.BodyParser(&t)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	// Regra id valido
	isValidId, err := strconv.Atoi(id)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	if isValidId > 5 {
		return c.SendStatus(fiber.StatusNotFound)
	}
	// Regra inteiro positivo
	if t.Valor <= 0 || t.Descricao == "" || len(t.Descricao) > 10 {
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
	cliente := domain.ClienteDomain{
		Id: id,
	}
	clientDomainResult := td.repository.GetClient(cliente)

	if clientDomainResult == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	switch strings.ToLower(t.Tipo) {
	case "c":
		transactionDomain := domain.TransactionDomain{
			UserId:    id,
			Valor:     t.Valor,
			Tipo:      t.Tipo,
			Descricao: t.Descricao,
		}
		clientDomainResult.Saldo += t.Valor

		td.repository.Insert(transactionDomain)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"limite": clientDomainResult.Limite, "saldo": clientDomainResult.Saldo})

	case "d":
		transactionDomain := domain.TransactionDomain{
			UserId:    id,
			Valor:     t.Valor,
			Tipo:      t.Tipo,
			Descricao: t.Descricao,
		}

		td.repository.Insert(transactionDomain)
		saldoFinal := clientDomainResult.Saldo - t.Valor
		semLimite := (clientDomainResult.Limite + saldoFinal) < 0
		if semLimite {
			return c.SendStatus(fiber.StatusUnprocessableEntity)
		}

		clientDomainResult.Saldo -= t.Valor

		td.repository.Update(*clientDomainResult)
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"limite": clientDomainResult.Limite, "saldo": clientDomainResult.Saldo})

	default:
		return c.SendStatus(fiber.StatusUnprocessableEntity)
	}
}
