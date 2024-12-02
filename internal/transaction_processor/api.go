package transaction_processor

import (
	"context"
)

type TransactionProcessor interface {
	Process(ctx context.Context) error

	Wait()
}
