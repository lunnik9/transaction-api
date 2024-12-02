package main

import (
	"context"
	"local/transaction/config"
	controller "local/transaction/internal/controller/http"
	"local/transaction/internal/domain"
	"local/transaction/internal/external/transaction_fetcher/ethereum"
	offset_manager "local/transaction/internal/repository/offset_manager/memory"
	transactions_repository "local/transaction/internal/repository/transactions/memory"
	"local/transaction/internal/rpc/client"
	"local/transaction/internal/service/parser"
	memory_storage "local/transaction/internal/storage/memory"
	"local/transaction/internal/transaction_processor/processor"
	"local/transaction/internal/worker"
	"log"
)

func main() {
	ctx := context.Background()

	path, err := config.ParseFlags()
	if err != nil {
		log.Fatalf("can't parse path flag %s", err)
	}

	config, err := config.NewConfig(path)
	if err != nil {
		log.Fatalf("can't create config %s", err)
	}

	ethereumConn := client.NewClientWithOpts(config.EthereumUrl, &client.RPCClientOpts{})
	ethereumCLient := ethereum.New(ethereumConn)

	offsetStorage := memory_storage.New[string, string]()
	offsetManager := offset_manager.New(offsetStorage)

	transactionsStorage := memory_storage.New[string, domain.Transaction]()
	transactionsRepo := transactions_repository.New(transactionsStorage)

	transactionProcessor := processor.New(ethereumCLient, offsetManager, transactionsRepo)

	service := parser.New(offsetManager, transactionsRepo)

	wrkr := worker.New(ctx, transactionProcessor, worker.WithTickerDuration(config.WorkerTickerDuration))
	wrkr.Start()

	ctrl := controller.NewTransactionController(service)

	err = ctrl.Run(config.Server.Port)
	if err != nil {
		log.Fatalf("cant run parser service controller")
	}

	transactionProcessor.Wait()
	wrkr.Stop()

}
