package mocks

import "climbingStuff/config"

type MockConfigProvider struct {
	Config config.Config
}

func NewMockConfigProvider(config config.Config) *MockConfigProvider {
	return &MockConfigProvider{Config: config}
}

func (m *MockConfigProvider) GetConfig() config.Config {
	return m.Config
}
