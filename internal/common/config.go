package common

import (
	"os"
)

type PostgresDatabaseConfig struct {
	ConnectionString string
}

func LoadPostgresDatabaseConfig() PostgresDatabaseConfig {

	// hostname := getEnv("hostname", "localhost")
	// hostport := getEnv("hostport ", 5432)
	// username := getEnv("host", "postgres")
	// password := getEnv("password", "kerala@21")
	// databasename := getEnv("databasename", "jobreport")
	// pgConString := fmt.Sprintf("port=%d host=%s user=%s password=%s dbname=%s sslmode=disable", hostport, hostname, username, password, databasename)

	return PostgresDatabaseConfig{
		ConnectionString: "postgres://odulyljiqysvkz:b6e71d857b6e35c6d9089ec32d14db79823cf36c4c9a52498fcd04791db51ee4@ec2-3-222-30-53.compute-1.amazonaws.com:5432/duas8p6afpl3r",
	}
}

func getEnv(env string, fallback interface{}) interface{} {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}
