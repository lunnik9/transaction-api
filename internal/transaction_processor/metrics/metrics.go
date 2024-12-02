package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	TransactionProcessorSavingErrorCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "transaction_processor_saving_error_counter",
			Help: "Transaction processor error happened",
		},
	)
)

func init() {
	prometheus.MustRegister(
		TransactionProcessorSavingErrorCounter,
	)
}
