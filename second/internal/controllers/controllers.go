package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"secondTask/internal/client"
	"sync"
	"time"
)

type ApiController struct {
	Client client.Client
}

func NewApiController(client client.Client) *ApiController {
	return &ApiController{Client: client}
}

var (
	currencyData client.CoinGeckoResponse
	lastUpdated  time.Time
	mu           sync.Mutex
)

func (a *ApiController) GetCryptoCurrencyPrice(ctx *gin.Context) {
	symbol := ctx.Param("symbol")

	if symbol == "" {
		log.Error().Msg("Symbol not provided")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Symbol not provided"})
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if currencyData == nil || time.Since(lastUpdated) > 10*time.Minute {
		log.Info().Msg("Fetching updated data from CoinGecko")
		currencyData = a.Client.FetchCoinGeckoData()
	}

	for _, currency := range currencyData {
		if currency.Symbol == symbol {
			log.Info().Msgf("Currency found: Symbol=%s, Price=%.2f", currency.Symbol, currency.Price)
			ctx.JSON(http.StatusOK, gin.H{"symbol": currency.Symbol, "price": currency.Price})
			return
		}
	}

	log.Warn().Msgf("Currency not found: Symbol=%s", symbol)
	ctx.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("Cryptocurrency with symbol %s not found", symbol)})
}
