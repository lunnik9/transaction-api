package memory

import (
	"context"
	"testing"

	"local/transaction/internal/domain"
	"local/transaction/internal/storage/memory"

	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	memStorage := memory.New[string, domain.Transaction]()
	repo := New(memStorage)
	ctx := context.Background()

	t.Run("AddSubscriber", func(t *testing.T) {
		err := repo.AddSubscriber(ctx, "subscriber1")
		assert.NoError(t, err, "Expected no error when adding a subscriber")

		addresses, err := repo.GetSubscriberAddresses(ctx)
		assert.NoError(t, err, "Expected no error when retrieving subscriber addresses")
		assert.Contains(t, addresses, "subscriber1", "Expected subscriber1 to be present in addresses")
	})

	t.Run("GetSubscriberAddresses", func(t *testing.T) {
		err := repo.AddSubscriber(ctx, "subscriber2")
		assert.NoError(t, err, "Expected no error when adding another subscriber")

		addresses, err := repo.GetSubscriberAddresses(ctx)
		assert.NoError(t, err, "Expected no error when retrieving subscriber addresses")
		assert.ElementsMatch(t, addresses, []string{"subscriber1", "subscriber2"}, "Expected both subscribers to be present")
	})

	t.Run("AddSubscriberTransactions", func(t *testing.T) {
		transactions := map[string][]domain.Transaction{
			"subscriber1": {
				{Hash: "tx1", From: "address1", To: "address2"},
			},
			"subscriber2": {
				{Hash: "tx2", From: "address3", To: "address4"},
			},
		}
		err := repo.AddSubscriberTransactions(ctx, transactions)
		assert.NoError(t, err, "Expected no error when adding transactions")

		tx, err := repo.GetSubscriberTransactions(ctx, "subscriber1")
		assert.NoError(t, err, "Expected no error when retrieving transactions for subscriber1")
		assert.Len(t, tx, 1, "Expected one transaction for subscriber1")
		assert.Equal(t, "tx1", tx[0].Hash, "Expected transaction hash to match for subscriber1")
	})

	t.Run("GetSubscriberTransactions for non-existing subscriber", func(t *testing.T) {
		tx, err := repo.GetSubscriberTransactions(ctx, "unknown_subscriber")
		assert.Error(t, err, "Expected an error for non-existing subscriber")
		assert.Nil(t, tx, "Expected transactions to be nil for non-existing subscriber")
	})
}
