package pokeapi

import (
	"net/http"
	"time"
)

type APIClient struct {
	Client http.Client
}

func NewClient(timeout time.Duration) APIClient {
	return APIClient{
		Client: http.Client{
			Timeout: timeout,
		},
	}
}