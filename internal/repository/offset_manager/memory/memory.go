package memory

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"local/transaction/internal/repository/offset_manager"
	"local/transaction/internal/storage"
)

const (
	blockOffsetKey = "block_offset"
)

type OffsetManager struct {
	memoryStorage storage.Storage[string, string]
}

func New(memoryStorage storage.Storage[string, string]) *OffsetManager {
	return &OffsetManager{memoryStorage: memoryStorage}
}

func (r *OffsetManager) GetOffset(ctx context.Context) (string, error) {
	vals, err := r.memoryStorage.Get(ctx, blockOffsetKey)
	if errors.Is(err, storage.ErrKeyNotFound) {
		return offset_manager.LatestBlock, nil
	}
	if err != nil {
		return "", err
	}

	if len(vals) != 1 {
		return "", fmt.Errorf("invalid length of stored values for block offset %d", len(vals))
	}

	return vals[0], nil
}

func (r *OffsetManager) SetNext(ctx context.Context, blockID string) error {
	blockNumber, err := hexToInt(blockID)
	if err != nil {
		return err
	}

	blockNumber++

	nextBlock := intToHex(blockNumber)

	return r.memoryStorage.Set(ctx, blockOffsetKey, []string{nextBlock})
}

func (r *OffsetManager) GetProcessed(ctx context.Context) (string, error) {
	currentOffset, err := r.GetOffset(ctx)
	if err != nil {
		return "", err
	}

	blockNumber, err := hexToInt(currentOffset)
	if err != nil {
		return "", err
	}

	blockNumber--

	return intToHex(blockNumber), nil
}

func hexToInt(blockID string) (int64, error) {
	if len(blockID) < 3 {
		return 0, fmt.Errorf("invalid block number")
	}

	blockNumber, err := strconv.ParseInt(blockID[2:], 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error converting block id: %s, err: %w", blockID, err)
	}

	return blockNumber, nil
}

func intToHex(blockNumber int64) string {
	return fmt.Sprintf("0x%x", blockNumber)
}
