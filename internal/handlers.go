package internal

import (
	"context"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	conn  = NewConn()
)

type Conta struct {
	Total       int    `json:"total"`
	Limite      int    `json:"limite"`
	DataExtrato string `json:"data_extrato"`
}

type Transacao struct {
	Valor     int    `json:"valor"`
	Tipo      string `json:"tipo"`
	Descricao string `json:"descricao"`
}

type Extrato struct {
	Saldo struct {
		Total       int    `json:"total"`
		Limite      int    `json:"limite"`
		DataExtrato string `json:"data_extrato"`
	} `json:"saldo"`
	UltimasTransacoes []UltimasTransacoes `json:"ultimas_transacoes"`
}

type UltimasTransacoes struct {
	Valor       int    `json:"valor"`
	Tipo        string `json:"tipo"`
	Descricao   string `json:"descricao"`
	RealizadaEm string `json:"realizada_em"`
}

func TransacaoHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	var t Transacao
	err := c.BodyParser(&t)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	// Regra id valido
	isValidId, err := strconv.Atoi(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}
	
	if isValidId > 5 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}
	// Regra inteiro positivo
	if t.Valor <= 0 || t.Descricao == "" || len(t.Descricao) > 10 {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

	// busca usuario
	filter := bson.M{"user_id": id}

	cliente := Clientes{}
	err = conn.Collection("clientes").FindOne(context.Background(), filter).Decode(&cliente)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	switch strings.ToLower(t.Tipo) {
	case "c":
		transacoes := Transacoes{
			UserId:      id,
			Valor:       t.Valor,
			Tipo:        t.Tipo,
			Descricao:   t.Descricao,
			RealizadaEm: primitive.NewDateTimeFromTime(time.Now()),
		}

		_, errT := conn.Collection("transacoes").InsertOne(context.Background(), &transacoes)
		if errT != nil {
			log.Println(err)
		}

		cliente.Saldo += t.Valor

		_, err := conn.Collection("clientes").UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{"saldo": cliente.Saldo}})
		if err != nil {
			log.Println(err)
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"limite": cliente.Limite, "saldo": cliente.Saldo})
	case "d":
		transacoes := Transacoes{
			UserId:      id,
			Valor:       t.Valor,
			Tipo:        t.Tipo,
			Descricao:   t.Descricao,
			RealizadaEm: primitive.NewDateTimeFromTime(time.Now()),
		}

		_, errT := conn.Collection("transacoes").InsertOne(context.Background(), &transacoes)
		if errT != nil {
			log.Println(err)

		}
		saldoFinal := cliente.Saldo - t.Valor
		semLimite := (cliente.Limite + saldoFinal) < 0
		if semLimite {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{})
		}

		cliente.Saldo -= t.Valor
		_, err := conn.Collection("clientes").UpdateOne(context.Background(), filter, bson.M{"$set": bson.M{
			"saldo": cliente.Saldo,
		}})
		if err != nil {
			log.Println(err)
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"limite": cliente.Limite, "saldo": cliente.Saldo})
	default:
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{})
	}

}

func ExtratoHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	// busca conta usuario
	filter := bson.D{{"user_id", id}}

	cliente := Clientes{}
	err := conn.Collection("clientes").FindOne(context.Background(), filter).Decode(&cliente)
	if err != nil {
		log.Println(err)
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{})
	}

	options := options.Find()
	options.SetSort(map[string]int{"realizada_em": -1})
	options.SetLimit(10)
	ctx := context.Background()
	cursor, err := conn.Collection("transacoes").Find(ctx, bson.M{}, options)
	if err != nil {
		log.Println(err)
	}

	defer cursor.Close(ctx)

	transacoes := []UltimasTransacoes{}
	for cursor.Next(ctx) {
		var transacao Transacoes
		if err := cursor.Decode(&transacao); err != nil {
			log.Println(err)
			continue
		}

		transacoes = append(transacoes, UltimasTransacoes{
			Valor:       transacao.Valor,
			Tipo:        transacao.Tipo,
			Descricao:   transacao.Descricao,
			RealizadaEm: time.Unix(int64(transacao.RealizadaEm), 0).Format("2006-01-02T15:04:05.999999Z"),
		})
	}

	// Verifique se houve algum erro durante a iteração do cursor
	if err := cursor.Err(); err != nil {
		log.Println(err)
	}

	extrato := Extrato{}
	extrato.Saldo.Limite = cliente.Limite
	extrato.Saldo.Total = cliente.Saldo
	extrato.Saldo.DataExtrato = time.Now().Format("2006-01-02T15:04:05.999999Z")
	extrato.UltimasTransacoes = transacoes

	// retornar extrato cliente 1
	return c.Status(fiber.StatusOK).JSON(extrato)

}
