package common

import (
	"fmt"
	"os"
)

type PostgresDatabaseConfig struct {
	ConnectionString string
}

func LoadPostgresDatabaseConfig() PostgresDatabaseConfig {

	hostname := getEnv("hostname", "localhost")
	hostport := getEnv("hostport ", 5432)
	username := getEnv("host", "postgres")
	password := getEnv("password", "kerala@21")
	databasename := getEnv("databasename", "postgres")
	pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", hostport, hostname, username, password, databasename)

	return PostgresDatabaseConfig{
		ConnectionString: pgConString,
	}
}

func getEnv(env string, fallback interface{}) interface{} {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
