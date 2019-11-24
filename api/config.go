package api

import (
	"os"

	"github.com/BurntSushi/toml"
)

var envVariables = []string{"MARIADB_SERVICE_HOST", "WORKDAY_DB_NAME", "WORKDAY_DB_USER", "WORKDAY_DB_PASSWORD"}

// DBConfig Represents database parameters
type DBConfig struct {
	Host     string
	Database string
	User     string
	Password string
}

func envVars(vars []string) bool {
	for _, v := range vars {
		if _, ok := os.LookupEnv(v); !ok {
			return false
		}
	}

	return true
}

// LoadDBConfig Read and parse DB configuration; environment variables have precedence over config file.
func LoadDBConfig(filepath string) (DBConfig, error) {
	var c DBConfig

	if envVars(envVariables) {
		host, _ := os.LookupEnv("MARIADB_SERVICE_HOST")
		name, _ := os.LookupEnv("WORKDAY_DB_NAME")
		user, _ := os.LookupEnv("WORKDAY_DB_USER")
		password, _ := os.LookupEnv("WORKDAY_DB_PASSWORD")

		c = DBConfig{host, name, user, password}
	} else {
		if _, err := toml.DecodeFile(filepath, &c); err != nil {
			return DBConfig{}, err
		}
	}

	return c, nil
}
