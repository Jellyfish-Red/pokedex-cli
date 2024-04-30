package pokeapi

import (
	"net/http"
	"time"

	"github.com/jellyfish-red/pokedex-cli/internal/pokecache"
)

type APIClient struct {
	Client http.Client
	LocationCache pokecache.APICache
}

func NewClient(timeout time.Duration) APIClient {
	return APIClient{
		Client: http.Client{
			Timeout: timeout,
		},
		LocationCache: pokecache.NewCache(5 * time.Second),
	}
}