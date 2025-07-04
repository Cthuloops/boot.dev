package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var ErrNoPreviousPage = errors.New("no previous page")

type Config struct {
	mapCall   int
	Next      string       `json:"next"`
	Previous  string       `json:"previous"`
	Locations [20]Location `json:"results"`
}

type Location struct {
	Name string
}

func (c *Config) GetNextPage() error {
	c.mapCall++
	if err := pokeApiRequest(c.Next, c); err != nil {
		return fmt.Errorf("error getting next page: %w", err)
	}
	return nil
}

func (c *Config) GetPreviousPage() error {
	if c.mapCall > 1 {
		c.mapCall--
	} else if c.mapCall == 1 {
		return ErrNoPreviousPage
	}
	if err := pokeApiRequest(c.Previous, c); err != nil {
		return fmt.Errorf("error getting previous page: %w", err)
	}
	return nil
}

func (c *Config) populateConfig() error {
	if err := pokeApiRequest("", c); err != nil {
		return fmt.Errorf("error populating config: %w", err)
	}
	return nil
}

func NewConfig() (*Config, error) {
	var config Config
	if err := config.populateConfig(); err != nil {
		return &Config{}, err
	}
	return &config, nil
}

func pokeApiRequest(url string, c *Config) error {
	if url == "" {
		url = "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	}
	res, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode > 299 {
		return fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("error reading result body: %w", err)
	}

	if err = json.Unmarshal(body, c); err != nil {
		return fmt.Errorf("error parsing json: %w", err)
	}

	return nil
}
