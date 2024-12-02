package controller

import (
	"local/transaction/internal/controller/metrics"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"local/transaction/internal/domain"
	"local/transaction/internal/service"
)

type TransactionController struct {
	parserService service.Parser
	router        *gin.Engine
}

func NewTransactionController(service service.Parser) *TransactionController {
	c := &TransactionController{
		parserService: service,
		router:        gin.New(),
	}

	c.router.Use(gin.Logger())
	c.router.Use(gin.Recovery())

	c.router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	c.router.GET("/v1/subscriber/transactions/:address", c.GetTransactions)

	c.router.POST("/v1/subscriber/subscribe", c.Subscribe)

	c.router.GET("/v1/transactions/current", c.GetCurrentBlock)

	return c
}

func (c *TransactionController) Run(port string) error {
	return c.router.Run(":" + port)
}

func (c *TransactionController) GetTransactions(ctx *gin.Context) {
	now := time.Now()
	defer metrics.MethodRequestLatency.WithLabelValues("GetTransactions").Observe(metrics.SinceSeconds(now))
	metrics.MethodRequestCount.WithLabelValues("GetTransactions").Inc()

	val := ctx.GetString("address")
	if val == "" {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	transactions, err := c.parserService.GetTransactions(ctx, val)
	if err != nil {
		metrics.MethodErrorCount.WithLabelValues("GetTransactions").Inc()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, transactions)
}

func (c *TransactionController) Subscribe(ctx *gin.Context) {
	now := time.Now()
	defer metrics.MethodRequestLatency.WithLabelValues("Subscribe").Observe(metrics.SinceSeconds(now))

	req := domain.SubscribeRequest{}
	metrics.MethodRequestCount.WithLabelValues("Subscribe").Inc()

	err := ctx.BindJSON(&req)
	if err != nil || req.Address == "" {
		metrics.MethodErrorCount.WithLabelValues("Subscribe").Inc()
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = c.parserService.Subscribe(ctx, req.Address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.Status(http.StatusOK)
}

func (c *TransactionController) GetCurrentBlock(ctx *gin.Context) {
	now := time.Now()
	defer metrics.MethodRequestLatency.WithLabelValues("GetCurrentBlock").Observe(metrics.SinceSeconds(now))

	metrics.MethodRequestCount.WithLabelValues("GetCurrentBlock").Inc()

	blockID, err := c.parserService.GetCurrentBlock(ctx)
	if err != nil {
		metrics.MethodErrorCount.WithLabelValues("GetCurrentBlock").Inc()
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, domain.GetCurrentBlockResponse{
		CurrentBlock: blockID,
	})
}
