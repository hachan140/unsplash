package config

import (
	"github.com/kelseyhightower/envconfig"
	"log"
)

type UnsplashConfig struct {
	APIKey string `envconfig:"UNSPLASH_API_KEY"`
}

func NewUnsplashConfig() UnsplashConfig {
	var unsplashConfig UnsplashConfig
	if err := envconfig.Process("", &unsplashConfig); err != nil {
		log.Fatal("error when parsing UnsplashConfig, error:", err)
	}
	return unsplashConfig
}
