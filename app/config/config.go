package config

import "os"

func Load() *Config {
	return &Config{
		AppName: getEnv("APP_NAME", "fiber-api"),
		Port:    getEnv("APP_PORT", "8080"),
		Env:     getEnv("APP_ENV", "local"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
