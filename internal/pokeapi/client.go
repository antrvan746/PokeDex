package pokeapi

import (
	"net/http"
	"time"
	"github.com/antrvan746/pokedexcli/internal/pokecache"
)


// Client
type Client struct {
	httpClient http.Client
	cache pokecache.Cache
}

// New Client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client {
		httpClient: http.Client{
			Timeout: timeout,
		},
		cache: pokecache.NewCache(cacheInterval),
	}
}