package config

import (
	"pokego/internal/pokecache"
)

type Config struct {
	NextLocationsURL *string
	PrevLocationsURL *string
	Cache            *pokecache.Cache
}
