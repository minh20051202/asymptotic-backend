package ledger

import "github.com/minh20051202/ticket-system-backend/internal/shared"

type LedgerService interface {
	Charge(transaction *shared.Transaction) (*shared.Transaction, error)
	Deposit(transaction *shared.Transaction) (*shared.Transaction, error)

	GetAllTransactions() ([]*shared.Transaction, error)
}

type service struct {
	repo LedgerRepository
}

func NewService(repo LedgerRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Charge(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := s.repo.Charge(transaction)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *service) Deposit(transaction *shared.Transaction) (*shared.Transaction, error) {
	tx, err := s.repo.Deposit(transaction)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *service) GetAllTransactions() ([]*shared.Transaction, error) {
	txs, err := s.repo.GetAllTransactions()

	if err != nil {
		return nil, err
	}

	return txs, nil
}
