package parser

import (
	"context"
	"local/transaction/internal/domain"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	offset_manager_mock "local/transaction/internal/repository/offset_manager/mock"
	transactions_repo_mock "local/transaction/internal/repository/transactions/mock"
)

const latestBlockID = ""

func TestGetCurrentBlock(t *testing.T) {
	ctx := context.Background()
	ctrl := gomock.NewController(t)

	offsetManager := offset_manager_mock.NewMockBlockOffsetManager(ctrl)
	transactionsRepo := transactions_repo_mock.NewMockRepository(ctrl)

	parser := New(offsetManager, transactionsRepo)

	offsetManager.EXPECT().GetOffset(ctx).Return(latestBlockID, nil)

	blockID, err := parser.GetCurrentBlock(ctx)
	require.NoError(t, err)
	require.Equal(t, latestBlockID, blockID)
}

func TestSubscribe(t *testing.T) {
	var address = "some_address"

	ctx := context.Background()
	ctrl := gomock.NewController(t)

	offsetManager := offset_manager_mock.NewMockBlockOffsetManager(ctrl)
	transactionsRepo := transactions_repo_mock.NewMockRepository(ctrl)

	parser := New(offsetManager, transactionsRepo)

	transactionsRepo.EXPECT().AddSubscriber(ctx, address).Return(nil)

	err := parser.Subscribe(ctx, address)
	require.NoError(t, err)
}

func TestGetTransactions(t *testing.T) {
	var (
		address = "some_address"

		transactions = []domain.Transaction{{BlockHash: "some_hash"}}
	)

	ctx := context.Background()
	ctrl := gomock.NewController(t)

	offsetManager := offset_manager_mock.NewMockBlockOffsetManager(ctrl)
	transactionsRepo := transactions_repo_mock.NewMockRepository(ctrl)

	parser := New(offsetManager, transactionsRepo)

	transactionsRepo.EXPECT().GetSubscriberTransactions(ctx, address).Return(transactions, nil)

	res, err := parser.GetTransactions(ctx, address)
	require.NoError(t, err)
	require.Equal(t, transactions, res)
}
