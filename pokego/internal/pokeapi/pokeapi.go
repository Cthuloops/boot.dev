package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"pokego/internal/pokecache"
)

const (
	baseURL = "https://pokeapi.co/api/v2"
)

type Response struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"string"`
	} `json:"results"`
}

func PokeApiRequest(pageURL *string, cache *pokecache.Cache) (Response, error) {
	url := baseURL + "/location-area?offset=0&limit=20"
	if pageURL != nil {
		url = *pageURL
	}

	// If the entry is in the cache already
	if response, ok := cache.Get(url); ok {
		locations := Response{}
		json.Unmarshal(response, &locations)
		log.Printf("Cache entry was found; logging from ApiRequest\n")
		return locations, nil
	}

	res, err := http.Get(url)
	if err != nil {
		return Response{}, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Response{}, fmt.Errorf("error reading result body: %w", err)
	}

	locations := Response{}
	if err = json.Unmarshal(body, &locations); err != nil {
		return Response{}, fmt.Errorf("error parsing json: %w", err)
	}

	// Add new entry to the cache
	cache.Add(url, body)

	return locations, nil
}
