package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

const (
	baseURL      = "https://pokeapi.co/api/v2"
	locationArea = "/location-area/"
)

type PokeAPITypes interface {
	Locations | LocationArea
}

type Locations struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"string"`
	} `json:"results"`
}

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func pokeAPIGet[T PokeAPITypes](client *Client, pageURL *string) (T, error) {
	// Create the result struct
	var result T

	// Build the URL
	url := baseURL + locationArea
	if pageURL != nil {
		url += *pageURL
	}

	// Check if already stored in the cache
	if val, ok := client.cache.Get(url); ok {
		if err := json.Unmarshal(val, &result); err != nil {
			return result, err
		}
		return result, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return result, err
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return result, nil
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		return result, nil
	}

	client.cache.Add(url, data)
	return result, nil
}
