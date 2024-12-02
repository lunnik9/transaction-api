package service

import (
	"context"
	"local/transaction/internal/domain"
)

type Parser interface {
	GetCurrentBlock(ctx context.Context) (string, error)
	Subscribe(ctx context.Context, address string) error
	GetTransactions(ctx context.Context, address string) ([]domain.Transaction, error)
}
