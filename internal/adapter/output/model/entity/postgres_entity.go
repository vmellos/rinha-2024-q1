package entity

type Customer struct {
	Id             int
	Limit, Balance string
}

type Transaction struct {
	ID          string
	Valor       int
	UserId      string
	Tipo        string
	Descricao   string
	RealizadaEm string
}
