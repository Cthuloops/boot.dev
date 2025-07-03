package services

import (
	"strings"
	"testing"
)

func TestConfig(t *testing.T) {
	url := "https://pokeapi.co/api/v2/location-area/?offset=0&limit=20"
	config, err := NewConfig()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	secondPageUrl := strings.Replace(url, "0", "20", 1)
	if config.Next != secondPageUrl {
		t.Errorf("Incorrect next page %s;\n expected %s", config.Next, secondPageUrl)
	}
	if config.Previous != "" {
		t.Errorf("Incorrect previous page %s;\n expected \"\"", config.Previous)
	}

	config.GetNextPage()
	thirdPageUrl := strings.Replace(secondPageUrl, "20", "40", 1)
	if config.Next != thirdPageUrl {
		t.Errorf("Incorrect next page %s;\n expected %s", config.Next, thirdPageUrl)
	}
	if config.Previous != url {
		t.Errorf("Incorrect previous page %s;\n expected %s", config.Previous, url)
	}
}
