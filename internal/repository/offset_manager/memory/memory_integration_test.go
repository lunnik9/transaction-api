package memory

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"

	"local/transaction/internal/repository/offset_manager"
	"local/transaction/internal/storage/memory"

	"github.com/stretchr/testify/assert"
)

func TestOffsetManager(t *testing.T) {
	memStorage := memory.New[string, string]()
	offsetMgr := New(memStorage)
	ctx := context.Background()

	t.Run("GetOffset when no offset is set", func(t *testing.T) {
		offset, err := offsetMgr.GetOffset(ctx)
		assert.NoError(t, err, "Expected no error when offset is not set")
		assert.Equal(t, offset_manager.LatestBlock, offset, "Expected GetOffset to return LatestBlock when no offset is set")
	})

	t.Run("SetNext and GetOffset", func(t *testing.T) {
		err := offsetMgr.SetNext(ctx, "0x1a3")
		assert.NoError(t, err, "Expected no error when setting a new offset")

		offset, err := offsetMgr.GetOffset(ctx)
		assert.NoError(t, err, "Expected no error when getting the updated offset")
		assert.Equal(t, "0x1a4", offset, "Expected GetOffset to return the incremented block offset")
	})

	t.Run("SetNext with invalid blockID", func(t *testing.T) {
		err := offsetMgr.SetNext(ctx, "xxx")
		require.Error(t, err, "Expected an error when setting an invalid block ID")
		assert.Contains(t, err.Error(), "error converting block id", "Expected error message to indicate invalid block number")
	})

	t.Run("SetNext and overwrite", func(t *testing.T) {
		err := offsetMgr.SetNext(ctx, "0x1a5")
		assert.NoError(t, err, "Expected no error when overwriting the offset")

		offset, err := offsetMgr.GetOffset(ctx)
		assert.NoError(t, err, "Expected no error when getting the updated offset")
		assert.Equal(t, "0x1a6", offset, "Expected GetOffset to return the incremented block offset")
	})
}
