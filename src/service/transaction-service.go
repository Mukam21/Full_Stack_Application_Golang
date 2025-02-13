package service

import (
	"context"
	"full_stack_application/entity"
	"full_stack_application/repository"
)

type TransactionService interface {
	Add(ctx context.Context, transaction entity.Transaction) error
	Edit(ctx context.Context, transaction entity.Transaction) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
}

type transactionService struct {
	repository repository.TransactionRepository
}

func NewTransactionService(repository repository.TransactionRepository) TransactionService {
	return &transactionService{
		repository: repository,
	}
}

func (s *transactionService) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	transactions, err := s.repository.FindAll(ctx)
	return transactions, err
}

func (s *transactionService) Add(ctx context.Context, transaction entity.Transaction) error {
	err := s.repository.Add(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) Edit(ctx context.Context, transaction entity.Transaction) error {
	err := s.repository.Edit(ctx, transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *transactionService) Delete(ctx context.Context, id int) error {
	err := s.repository.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
