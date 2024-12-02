package transaction_fetcher

import (
	"context"
	"local/transaction/internal/domain"
)

type TransactionFetcher interface {
	GetBlockTransactions(ctx context.Context, blockID string) ([]domain.Transaction, error)
}
