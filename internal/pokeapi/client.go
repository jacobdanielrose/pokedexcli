package pokeapi

import (
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

// Client -
type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// NewClient -
func NewClient(timeout time.Duration, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
