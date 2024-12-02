package parser

import (
	"context"
	"log"

	"local/transaction/internal/domain"
	"local/transaction/internal/repository/offset_manager"
	transactions_repo "local/transaction/internal/repository/transactions"
)

type Parser struct {
	offsetManager    offset_manager.BlockOffsetManager
	transactionsRepo transactions_repo.Repository
}

func New(offsetManager offset_manager.BlockOffsetManager, transactionsRepo transactions_repo.Repository) *Parser {
	return &Parser{
		offsetManager:    offsetManager,
		transactionsRepo: transactionsRepo,
	}
}

func (p Parser) GetCurrentBlock(ctx context.Context) (string, error) {
	blockID, err := p.offsetManager.GetProcessed(ctx)
	if err != nil {
		log.Printf("error on GetCurrentBlock %s", err)
		return "", err
	}

	return blockID, nil
}

func (p Parser) Subscribe(ctx context.Context, address string) error {
	err := p.transactionsRepo.AddSubscriber(ctx, address)
	if err != nil {
		log.Printf("error on Subscribe %s", err)
		return err
	}

	return nil
}

func (p Parser) GetTransactions(ctx context.Context, address string) ([]domain.Transaction, error) {
	transactions, err := p.transactionsRepo.GetSubscriberTransactions(ctx, address)
	if err != nil {
		log.Printf("error on GetTransactions %s", err)
		return nil, err
	}

	return transactions, nil
}
