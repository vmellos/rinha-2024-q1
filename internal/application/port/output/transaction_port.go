package output

import "rinha_2024_q1/internal/application/domain"

type TransactionPort interface {
	Find(clientDomain domain.ClienteDomain) []*domain.TransactionDomain
	Insert(transactionDomain domain.TransactionDomain) *domain.TransactionDomain
	Update(transactionDomain domain.ClienteDomain) *domain.ClienteDomain
	GetClient(clientDomain domain.ClienteDomain) *domain.ClienteDomain
}
