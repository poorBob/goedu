package tests

import (
	"testing"

	"climbingStuff/config"
	"climbingStuff/tests/mocks"
)

func TestRunApplicationWithMockConfig(t *testing.T) {
	mockConfig := config.Config{
		Server:   "mockServer",
		Port:     1234,
		Database: "mockDB",
	}

	mockProvider := mocks.NewMockConfigProvider(mockConfig)

	cfg := mockProvider.GetConfig()

	if cfg.Server != "mockServer" {
		t.Errorf("Expected Server to be 'mockServer', got '%s'", cfg.Server)
	}

	if cfg.Port != 1234 {
		t.Errorf("Expected Server to be '1234', got '%v'", cfg.Port)
	}
}
