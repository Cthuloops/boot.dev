package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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

func PokeApiRequest(pageURL *string) (Response, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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

	return locations, nil
}
