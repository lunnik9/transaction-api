package processor

import (
	"context"
	"local/transaction/internal/external/transaction_fetcher"
	"log"
	"sync"

	"local/transaction/internal/domain"
	"local/transaction/internal/repository/offset_manager"
	transactions_repo "local/transaction/internal/repository/transactions"
	"local/transaction/internal/transaction_processor/metrics"
)

type TransactionProcessor struct {
	transactionsFetcher transaction_fetcher.TransactionFetcher
	blockOffsetManager  offset_manager.BlockOffsetManager
	subscribersRepo     transactions_repo.Repository

	wg *sync.WaitGroup
}

func New(transactionsFetcher transaction_fetcher.TransactionFetcher, blockOffsetManager offset_manager.BlockOffsetManager, subscribersRepo transactions_repo.Repository) *TransactionProcessor {
	return &TransactionProcessor{
		transactionsFetcher: transactionsFetcher,
		blockOffsetManager:  blockOffsetManager,
		subscribersRepo:     subscribersRepo,

		wg: &sync.WaitGroup{},
	}
}

func (p *TransactionProcessor) Process(ctx context.Context) error {
	blockOffset, err := p.blockOffsetManager.GetOffset(ctx)
	if err != nil {
		return err
	}

	transactions, err := p.transactionsFetcher.GetBlockTransactions(ctx, blockOffset)
	if err != nil {
		return err
	}

	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		saveErr := p.saveTransaction(ctx, transactions)
		if saveErr != nil {
			log.Printf("error on saving transactions: %s", saveErr)
			metrics.TransactionProcessorSavingErrorCounter.Inc()
		}
	}()

	if len(transactions) != 0 {
		return p.blockOffsetManager.SetNext(ctx, transactions[0].BlockNumber)
	}

	return nil
}

func (p *TransactionProcessor) saveTransaction(ctx context.Context, fetchedTransactions []domain.Transaction) error {
	subscribers, err := p.subscribersRepo.GetSubscriberAddresses(ctx)
	if err != nil {
		return err
	}

	subscribersTransactions := make(map[string][]domain.Transaction)
	for _, subscriber := range subscribers {
		subscribersTransactions[subscriber] = make([]domain.Transaction, 0)
	}

	for _, transaction := range fetchedTransactions {
		from, to := transaction.From, transaction.To

		if transactions, shouldCollect := subscribersTransactions[from]; shouldCollect {
			subscribersTransactions[from] = append(transactions, transaction)
		}

		if from == to {
			continue
		}

		if transactions, shouldCollect := subscribersTransactions[to]; shouldCollect {
			subscribersTransactions[to] = append(transactions, transaction)
		}
	}

	return p.subscribersRepo.AddSubscriberTransactions(ctx, subscribersTransactions)
}

func (p *TransactionProcessor) Wait() {
	p.wg.Wait()
}
