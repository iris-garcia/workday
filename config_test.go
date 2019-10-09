package workday

import (
	"os"
	"testing"
)

// Test envVars checks the environment variables correctly
func TestEnvVars(t *testing.T) {
	env_vars := []string{"VAR1", "VAR2"}
	os.Setenv("VAR1", "value")
	os.Setenv("VAR2", "value")

	if !envVars(env_vars) {
		t.Error("Env var not found and it was set.")
	}

	env_vars = []string{"VAR1", "VAR2", "VAR3"}

	if envVars(env_vars) {
		t.Error("Env var was not set and it is found.")
	}
}

// Test DB config is read correctly from a file
func TestDBConfigFile(t *testing.T) {
	cfg, err := LoadDBConfig("./db_config.toml")
	if err != nil {
		t.Errorf("Error loading db_config file: %v", err.Error())
	}

	if cfg.Host != "127.0.0.1" {
		t.Errorf("Wrong host parameter, got %v - expected %v", cfg.Host, "127.0.0.1")
	} else if cfg.Database != "workday" {
		t.Errorf("Wrong database parameter, got %v - expected %v", cfg.Database, "workday")
	} else if cfg.User != "workday" {
		t.Errorf("Wrong user parameter, got %v - expected %v", cfg.User, "workday")
	} else if cfg.Password != "changeme" {
		t.Errorf("Wrong password parameter, got %v - expected %v", cfg.Password, "changeme")
	}
}

// Test DB config is read correctly from a environment variables
func TestDBConfigEnv(t *testing.T) {
	os.Setenv("WORKDAY_DB_HOST", "1.1.1.1")
	os.Setenv("WORKDAY_DB_NAME", "name")
	os.Setenv("WORKDAY_DB_USER", "user")
	os.Setenv("WORKDAY_DB_PASSWORD", "pass")

	cfg, err := LoadDBConfig("./db_config.toml")
	if err != nil {
		t.Errorf("Error loading db_config file: %v", err.Error())
	}

	if cfg.Host != "1.1.1.1" {
		t.Errorf("Wrong host parameter, got %v - expected %v", cfg.Host, "1.1.1.1")
	} else if cfg.Database != "name" {
		t.Errorf("Wrong database parameter, got %v - expected %v", cfg.Database, "name")
	} else if cfg.User != "user" {
		t.Errorf("Wrong user parameter, got %v - expected %v", cfg.User, "user")
	} else if cfg.Password != "pass" {
		t.Errorf("Wrong password parameter, got %v - expected %v", cfg.Password, "pass")
	}
}
