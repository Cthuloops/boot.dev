package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Locations struct {
	Count     int     `json:"count"`
	Next      *string `json:"next"`
	Previous  *string `json:"previous"`
	Locations []struct {
		Name string `json:"name"`
		URL  string `json:"string"`
	} `json:"results"`
}

func (c *Client) ListLocations(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locations := Locations{}
		if err := json.Unmarshal(val, &locations); err != nil {
			return Locations{}, nil
		}
		return locations, nil
	}

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, nil
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return Locations{}, nil
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	c.cache.Add(url, data)
	return locations, nil
}
