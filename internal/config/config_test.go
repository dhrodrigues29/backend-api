package config

import "testing"

func TestLoadUsesDefaultServerAddress(t *testing.T) {
	t.Setenv("SERVER_ADDRESS", "")

	cfg := Load()

	if cfg.ServerAddress != ":8080" {
		t.Fatalf("expected default server address :8080, got %s", cfg.ServerAddress)
	}
}

func TestLoadUsesEnvironmentServerAddress(t *testing.T) {
	t.Setenv("SERVER_ADDRESS", ":9090")

	cfg := Load()

	if cfg.ServerAddress != ":9090" {
		t.Fatalf("expected server address :9090, got %s", cfg.ServerAddress)
	}
}