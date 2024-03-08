package request

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

