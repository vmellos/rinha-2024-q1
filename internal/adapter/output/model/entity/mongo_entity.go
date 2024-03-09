package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

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
