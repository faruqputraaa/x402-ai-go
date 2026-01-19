package config

import (
	"log"
	"os"
)

type Config struct {
	BackendAddress string
	BaseRPC        string
	PriceWei       int64
	GeminiAPIKey   string
}

func Load() Config {
	cfg := Config{
		BackendAddress: os.Getenv("BACKEND_ADDRESS"),
		BaseRPC:        os.Getenv("BASE_RPC"),
		GeminiAPIKey:   os.Getenv("GEMINI_API_KEY"),
		PriceWei:       100000000000, // 0.0000001 ETH
	}

	if cfg.BaseRPC == "" {
		log.Fatal("BASE_RPC is empty")
	}
	if cfg.BackendAddress == "" {
		log.Fatal("BACKEND_ADDRESS is empty")
	}
	if cfg.GeminiAPIKey == "" {
		log.Fatal("GEMINI_API_KEY is empty")
	}

	return cfg
}
