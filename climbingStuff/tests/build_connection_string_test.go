package tests

import (
	"climbingStuff/config"
	"climbingStuff/utils"
	"testing"
)

func TestBuildConnectionString(t *testing.T) {
	tests := []struct {
		name           string
		server         string
		port           int
		database       string
		expectedOutput string
	}{
		{
			name:           "Valid connection string",
			server:         "localhost",
			port:           1433,
			database:       "testDB",
			expectedOutput: "server=localhost;port=1433;database=testDB;trusted_connection=yes",
		},
		{
			name:           "Empty server",
			server:         "",
			port:           1433,
			database:       "testDB",
			expectedOutput: "server=;port=1433;database=testDB;trusted_connection=yes",
		},
		{
			name:           "Empty database",
			server:         "localhost",
			port:           1433,
			database:       "",
			expectedOutput: "server=localhost;port=1433;database=;trusted_connection=yes",
		},
		{
			name:           "Zero port",
			server:         "localhost",
			port:           0,
			database:       "testDB",
			expectedOutput: "server=localhost;port=0;database=testDB;trusted_connection=yes",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := utils.BuildConnectionString(config.Config{Server: tt.server, Port: tt.port, Database: tt.database})
			if output != tt.expectedOutput {
				t.Errorf("BuildConnectionString() = %v, want %v", output, tt.expectedOutput)
			}
		})
	}
}
