package config

import "os"

type Config struct {
	ServerAddress string
}

func Load() Config {
	serverAddress := os.Getenv("SERVER_ADDRESS")

	if serverAddress == "" {
		serverAddress = ":8080"
	}

	return Config{
		ServerAddress: serverAddress,
	}
}