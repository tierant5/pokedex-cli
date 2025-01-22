package pokeapi

import (
	"net/http"
	"time"

	"github.com/tierant5/pokedex-cli/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	client := Client{}
	client.httpClient = http.Client{
		Timeout: timeout,
	}
	client.cache = pokecache.NewCache(cacheInterval)
	return client
}
