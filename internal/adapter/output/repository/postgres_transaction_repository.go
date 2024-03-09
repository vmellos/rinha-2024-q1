package repository

import (
	"database/sql"
	"log"
	"rinha_2024_q1/internal/adapter/output/model/entity"
	"rinha_2024_q1/internal/application/domain"
	"rinha_2024_q1/internal/application/port/output"
	"strconv"
)

func NewTransactionPostgresRepository(database *sql.DB) output.TransactionPort {
	return &transactionPostgresRepository{database}
}

type transactionPostgresRepository struct {
	databaseConnection *sql.DB
}

func (tr *transactionPostgresRepository) Find(clientDomain domain.ClienteDomain) []*domain.TransactionDomain {
	id, _ := strconv.Atoi(clientDomain.Id)
	rows, err := tr.databaseConnection.Query(`select h.value, h.type, h.description, h.do_at from rinha.history h where h.user_id = $1 order by id desc limit 10`, id)
	if err != nil {
		log.Print(err)
		return []*domain.TransactionDomain{}
	}

	transactions := []*domain.TransactionDomain{}
	for rows.Next() {
		var transaction entity.Transaction
		err := rows.Scan(&transaction.Valor, &transaction.Tipo, &transaction.Descricao, &transaction.RealizadaEm)
		if err != nil {
			log.Print(err)
			return []*domain.TransactionDomain{}
		}
		transactions = append(transactions, &domain.TransactionDomain{
			Valor:     transaction.Valor,
			Descricao: transaction.Descricao,
		})
	}

	return transactions
}

func (tr *transactionPostgresRepository) Insert(transactionDomain domain.TransactionDomain) *domain.TransactionDomain {

	id, _ := strconv.Atoi(transactionDomain.UserId)
	_, err := tr.databaseConnection.Exec(`
	insert into rinha.history (value, type, description, user_id) values($1, $2, $3, $4)`, 
	transactionDomain.Valor, transactionDomain.Tipo, transactionDomain.Descricao, id)

	if err != nil {
		log.Print(err)
		return &domain.TransactionDomain{}
	}

	return &domain.TransactionDomain{}
}

func (tr *transactionPostgresRepository) Update(clienteDomain domain.ClienteDomain) *domain.ClienteDomain {
	id, _ := strconv.Atoi(clienteDomain.Id)
	_, err := tr.databaseConnection.Exec(`
	update rinha.users
	set initial_balance = $1 
	where id = $2`, clienteDomain.Saldo, id)

	if err != nil {
		log.Print(err)
		return &domain.ClienteDomain{}
	}

	return &domain.ClienteDomain{}
}

func (tr *transactionPostgresRepository) GetClient(clientDomain domain.ClienteDomain) *domain.ClienteDomain {
	rows, err := tr.databaseConnection.Query(`select * from rinha.users where id = $1`, clientDomain.Id)
	if err != nil {
		log.Print(err)
		return &domain.ClienteDomain{}
	}

	var c entity.Customer
	for rows.Next() {
		err := rows.Scan(&c.Id, &c.Limit, &c.Balance)
		if err != nil {
			log.Print(err)
			return &domain.ClienteDomain{}
		}
	}
	limit, _ := strconv.Atoi(c.Limit)
	balance, _ := strconv.Atoi(c.Balance)
	return &domain.ClienteDomain{
		Id:     strconv.Itoa(c.Id),
		Limite: limit,
		Saldo:  balance,
	}
}
