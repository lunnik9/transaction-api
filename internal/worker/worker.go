package worker

import (
	"context"
	"log"
	"time"

	"local/transaction/internal/transaction_processor"
)

type Option func(*options)

type options struct {
	tickerDuration time.Duration
}

type Worker struct {
	ctx context.Context

	cfg *options

	transactionProcessor transaction_processor.TransactionProcessor

	stopped chan struct{}
	started chan struct{}
	done    chan struct{}
}

func New(ctx context.Context,
	transactionProcessor transaction_processor.TransactionProcessor,
	opts ...Option) *Worker {
	worker := &Worker{
		ctx: ctx,

		transactionProcessor: transactionProcessor,

		stopped: make(chan struct{}),
		started: make(chan struct{}),
		done:    make(chan struct{}),
	}

	worker.cfg = &options{
		tickerDuration: time.Second,
	}

	for _, opt := range opts {
		opt(worker.cfg)
	}

	return worker
}

func (w *Worker) Start() {
	go w.loop()
	<-w.started
}

func (w *Worker) Stop() {
	close(w.done)
	<-w.stopped
}

func (w *Worker) loop() {
	t := time.NewTimer(w.cfg.tickerDuration)
	defer t.Stop()

	close(w.started)
	defer close(w.stopped)

	for {
		select {
		case <-w.done:
			return
		case <-t.C:
			t.Reset(w.cfg.tickerDuration)

			err := w.transactionProcessor.Process(w.ctx)
			if err != nil {
				log.Printf("worker get process error: %s", err.Error())
			}
		}
	}
}

func WithTickerDuration(duration time.Duration) Option {
	return func(o *options) {
		o.tickerDuration = duration
	}
}
