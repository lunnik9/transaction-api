package transactions

import (
	"context"
	"local/transaction/internal/domain"
)

//go:generate mockgen --build_flags=--mod=mod -destination ./mock/api.go -package mock local/transaction/internal/repository/transactions Repository

type Repository interface {
	AddSubscriber(ctx context.Context, subscriber string) error
	GetSubscriberAddresses(ctx context.Context) ([]string, error)
	AddSubscriberTransactions(ctx context.Context, subscriberTransactions map[string][]domain.Transaction) error
	GetSubscriberTransactions(ctx context.Context, subscriber string) ([]domain.Transaction, error)
}
