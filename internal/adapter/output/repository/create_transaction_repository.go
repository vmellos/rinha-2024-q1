package repository

import (
	"context"
	"log"
	"rinha_2024_q1/internal/adapter/input/model/response"
	"rinha_2024_q1/internal/application/domain"
	"rinha_2024_q1/internal/application/port/output"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collection   *mongo.Collection
	DatabaseName = "rinha"
)

func NewTransactionRepository(database *mongo.Database) output.TransactionPort {
	return &transactionRepository{database}
}

type transactionRepository struct {
	databaseConnection *mongo.Database
}

func (tr *transactionRepository) Find(clientDomain domain.ClienteDomain) []*domain.TransactionDomain {
	options := options.Find()
	options.SetSort(map[string]int{"realizada_em": -1})
	options.SetLimit(10)
	ctx := context.Background()
	cursor, err := tr.databaseConnection.Collection("transacoes").Find(ctx, bson.M{}, options)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(ctx)

	transacoes := []response.UltimasTransacoes{}
	for cursor.Next(ctx) {
		var transacao response.Transacoes
		if err := cursor.Decode(&transacao); err != nil {
			log.Println(err)
			continue
		}

		transacoes = append(transacoes, response.UltimasTransacoes{
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

	extrato := response.Extrato{}
	extrato.Saldo.Limite = clientDomain.Limite
	extrato.Saldo.Total = clientDomain.Saldo
	extrato.Saldo.DataExtrato = time.Now().Format("2006-01-02T15:04:05.999999Z")
	extrato.UltimasTransacoes = transacoes

	// converter to domain
	transacoesDomain := []*domain.TransactionDomain{}
	for _, transacao := range transacoes {
		transacoesDomain = append(transacoesDomain, &domain.TransactionDomain{
			Valor:       transacao.Valor,
			Tipo:        transacao.Tipo,
			Descricao:   transacao.Descricao,
			RealizadaEm: transacao.RealizadaEm,
		})
	}

	return transacoesDomain
}

func (tr *transactionRepository) Insert(transactionDomain domain.TransactionDomain) *domain.TransactionDomain {

	transactionDomain.RealizadaEm = primitive.NewDateTimeFromTime(time.Now()).Time().Format("2006-01-02T15:04:05.999999Z")
	_, err := tr.databaseConnection.Collection("transacoes").InsertOne(context.Background(), &transactionDomain)
	if err != nil {
		log.Println(err)
		return &domain.TransactionDomain{}
	}
	return &transactionDomain
}

func (tr *transactionRepository) Update(transactionDomain domain.TransactionDomain) *domain.TransactionDomain {
	_ = tr.databaseConnection.Collection("clientes").FindOneAndUpdate(context.Background(), bson.M{"user_id": transactionDomain.UserId}, transactionDomain)
	return &transactionDomain
}

func (tr *transactionRepository) GetClient(clientDomain domain.ClienteDomain) *domain.ClienteDomain {
	filter := bson.D{{"user_id", clientDomain.Id}}
	_ = tr.databaseConnection.Collection("clientes").FindOne(context.Background(), filter).Decode(&clientDomain)
	return &clientDomain
}
