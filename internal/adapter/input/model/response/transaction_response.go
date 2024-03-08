package response

import "go.mongodb.org/mongo-driver/bson/primitive"

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

type Transacoes struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Valor       int                `bson:"valor"`
	UserId      string             `bson:"user_id"`
	Tipo        string             `bson:"tipo"`
	Descricao   string             `bson:"descricao"`
	RealizadaEm primitive.DateTime `bson:"realizada_em"`
}
