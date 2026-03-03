package ledger

import (
	"time"

	"github.com/google/uuid"
)

type LedgerService interface {
	ChargeWithRequest(req *CreateTransactionRequest) (*Transaction, error)
	Deposit(req *CreateTransactionRequest) (*Transaction, error)

	GetAllTransactions() ([]*Transaction, error)
	Charge(userId uuid.UUID, amount int64, idempotencyKey string) error
}

type service struct {
	repo LedgerRepository
}

func NewService(repo LedgerRepository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) Charge(userId uuid.UUID, amount int64, idempotencyKey string) error {
	req := &CreateTransactionRequest{
		UserId:         userId,
		Amount:         amount,
		IdempotencyKey: idempotencyKey,
		Type:           "CHARGE",
	}

	_, err := s.ChargeWithRequest(req)
	if err != nil {
		return err
	}
	return nil
}

func (s *service) ChargeWithRequest(req *CreateTransactionRequest) (*Transaction, error) {
	newTransaction := &Transaction{
		TransactionId:  uuid.New(),
		UserId:         req.UserId,
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Type:           "CHARGE",
		CreatedAt:      time.Now().UTC(),
	}
	tx, err := s.repo.Charge(newTransaction)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *service) Deposit(req *CreateTransactionRequest) (*Transaction, error) {
	newTransaction := &Transaction{
		TransactionId:  uuid.New(),
		UserId:         req.UserId,
		IdempotencyKey: req.IdempotencyKey,
		Amount:         req.Amount,
		Type:           "DEPOSIT",
		CreatedAt:      time.Now().UTC(),
	}
	tx, err := s.repo.Deposit(newTransaction)

	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (s *service) GetAllTransactions() ([]*Transaction, error) {
	txs, err := s.repo.GetAllTransactions()

	if err != nil {
		return nil, err
	}

	return txs, nil
}
