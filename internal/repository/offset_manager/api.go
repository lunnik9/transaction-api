package offset_manager

import "context"

//go:generate mockgen --build_flags=--mod=mod -destination ./mock/api.go -package mock local/transaction/internal/repository/offset_manager BlockOffsetManager

// BlockOffsetManager is an interface for managing the block offset used to fetch blocks.
// If no offset is set, it returns the constant `LatestBlock`.
// While the current implementation uses in-memory storage, it can be extended to use Redis
// or other persistent storage in the future to avoid losing the offset during service restarts.
type BlockOffsetManager interface {
	// GetOffset retrieves the current block offset.
	// If no offset is set, it should return the constant `LatestBlock`.
	GetOffset(ctx context.Context) (string, error)
	// SetNext updates the block offset to the provided block ID.
	SetNext(ctx context.Context, blockID string) error
	// GetProcessed - returns last processed transaction block ID
	GetProcessed(ctx context.Context) (string, error)
}

// LatestBlock represents the default value for the block offset
// when no offset has been explicitly set.
const LatestBlock = "latest"
