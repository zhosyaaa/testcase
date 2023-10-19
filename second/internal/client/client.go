package client

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"io/ioutil"
	"net/http"
	"secondTask/internal/models"
	"sync"
	"time"
)

type Client struct {
	apiURL string
}

func NewClient(apiURL string) *Client {
	return &Client{apiURL: apiURL}
}

type CoinGeckoResponse []models.Currency

var (
	currencyData CoinGeckoResponse
	lastUpdated  time.Time
	mu           sync.Mutex
)

func (c *Client) FetchCoinGeckoData() CoinGeckoResponse {
	resp, err := http.Get(c.apiURL)
	if err != nil {
		log.Error().Err(err).Msg("Error fetching data")
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Error().Err(err).Msg("Error reading data")
		return nil
	}

	var data CoinGeckoResponse
	if err := json.Unmarshal(body, &data); err != nil {
		log.Error().Err(err).Msg("Error parsing data")
		return nil
	}

	mu.Lock()
	currencyData = data
	lastUpdated = time.Now()
	mu.Unlock()
	return currencyData
}
