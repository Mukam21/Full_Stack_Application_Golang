package repository

import (
	"context"
	"full_stack_application/entity"

	"github.com/jmoiron/sqlx"
)

type TransactionRepository interface {
	Add(ctx context.Context, u entity.Transaction) error
	Edit(ctx context.Context, u entity.Transaction) error
	Delete(ctx context.Context, id int) error
	FindAll(ctx context.Context) ([]entity.Transaction, error)
}

type pgTransactionRepository struct {
	DB *sqlx.DB
}

func NewTransactionRepository(db *sqlx.DB) TransactionRepository {
	return &pgTransactionRepository{
		DB: db,
	}
}

func (r *pgTransactionRepository) FindAll(ctx context.Context) ([]entity.Transaction, error) {
	transactions := []entity.Transaction{}

	query := "SELECT * FROM transactions"

	if err := r.DB.Select(&transactions, query); err != nil {
		return nil, err
	}

	return transactions, nil
}

func (r *pgTransactionRepository) Add(ctx context.Context, u entity.Transaction) error {
	query := "INSERT INTO transactions (type, ticker, volume, price, date) VALUES ($1, $2, $3, $4, $5) RETURNING id"

	if err := r.DB.QueryRow(query, u.Type, u.Ticker, u.Volume, u.Price, u.Date).Scan(&u.ID); err != nil {
		return err
	}

	return nil
}

func (r *pgTransactionRepository) Edit(ctx context.Context, u entity.Transaction) error {
	query := "UPDATE transaction SET (type, ticker, volume, price, date) = ($1, $2, $3, $4, $5) WHERE id=$6"

	if _, err := r.DB.ExecContext(ctx, query, u.Type, u.Ticker, u.Volume, u.Price, u.Date, u.ID); err != nil {
		return err
	}

	return nil
}

func (r *pgTransactionRepository) Delete(ctx context.Context, id int) error {
	query := "DELETE FROM transactions WHERE id=$1"

	if _, err := r.DB.Exec(query, id); err != nil {
		return err
	}

	return nil
}
