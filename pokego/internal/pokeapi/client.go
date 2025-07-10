package pokeapi

import (
	"net/http"
	"time"

	"pokego/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

func (c *Client) ListLocations(location *string) (Locations, error) {
	return pokeAPIGet[Locations](c, location)
}

func (c *Client) PokemonAtLocation(location *string) (LocationArea, error) {
	return pokeAPIGet[LocationArea](c, location)
}

func (c *Client) PokemonInfo(pokemon *string) (Pokemon, error) {
	return pokeAPIGet[Pokemon](c, pokemon)
}
