package memory

import (
	"context"
	"local/transaction/internal/domain"
	"maps"
	"slices"

	"local/transaction/internal/storage"
)

type Repository struct {
	memoryStorage storage.Storage[string, domain.Transaction]
}

func New(memoryStorage storage.Storage[string, domain.Transaction]) *Repository {
	return &Repository{memoryStorage: memoryStorage}
}

func (r *Repository) AddSubscriber(ctx context.Context, subscriber string) error {
	return r.memoryStorage.Set(ctx, subscriber, []domain.Transaction{})
}

func (r *Repository) GetSubscriberAddresses(ctx context.Context) ([]string, error) {
	allSubscribers, err := r.memoryStorage.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	keys := slices.Collect(maps.Keys(allSubscribers))
	return keys, nil
}

func (r *Repository) AddSubscriberTransactions(ctx context.Context, subscriberTransactions map[string][]domain.Transaction) error {
	return r.memoryStorage.Append(ctx, subscriberTransactions)
}

func (r *Repository) GetSubscriberTransactions(ctx context.Context, subscriber string) ([]domain.Transaction, error) {
	return r.memoryStorage.Get(ctx, subscriber)
}
