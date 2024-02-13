package internal

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Collection   *mongo.Collection
	DatabaseName = "rinha"
)

type Transacoes struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Valor       int                `bson:"valor"`
	UserId      string             `bson:"user_id"`
	Tipo        string             `bson:"tipo"`
	Descricao   string             `bson:"descricao"`
	RealizadaEm primitive.DateTime `bson:"realizada_em"`
}

type Clientes struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Id     string             `bson:"user_id"`
	Limite int                `bson:"limite"`
	Saldo  int                `bson:"saldo"`
}

func NewConn(collectionName string) *mongo.Collection {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://rinha_db:27017"))

	if err != nil {
		log.Println(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Println(err)
	}

	Collection = client.Database(DatabaseName).Collection(collectionName)

	return Collection
}
