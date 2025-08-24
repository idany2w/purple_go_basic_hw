package config

import "os"

type Config struct {
	XMasterKey string
}

func NewConfig() *Config {
	XMasterKey := os.Getenv("X_MASTER_KEY")

	if XMasterKey == "" {
		panic("X_MASTER_KEY not found in ENV")
	}

	return &Config{
		XMasterKey: XMasterKey,
	}
}
