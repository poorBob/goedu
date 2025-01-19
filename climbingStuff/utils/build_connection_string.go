package utils

import (
	"climbingStuff/config"
	"fmt"
)

func BuildConnectionString(cfg config.Config) string {
	return fmt.Sprintf("server=%s;port=%d;database=%s;trusted_connection=yes", cfg.Server, cfg.Port, cfg.Database)
}
