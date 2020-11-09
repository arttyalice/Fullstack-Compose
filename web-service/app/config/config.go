package config

import (
	"fmt"
	"os"

	"github.com/Netflix/go-env"
)

type Config struct {
	ServePort string `env:"APP_SERVE_PORT"`

	// Postgresql
	PQHost     string `env:"APP_PQ_HOST"`
	PQPort     string `env:"APP_PQ_PORT"`
	PQUser     string `env:"APP_PQ_USER"`
	PQPassword string `env:"APP_PQ_PASSWORD"`
	PQDB       string `env:"APP_PQ_DB"`
}

func ParseConfig() (cfg *Config) {
	cfg = new(Config)
	fmt.Printf("Parsing ENV to application configuration...\n")
	fmt.Printf("Parsing v1 ENV...\n")
	_, err := env.UnmarshalFromEnviron(cfg)
	if err != nil {
		fmt.Printf("Cannot read config: %s\n", err.Error())
		os.Exit(1)
	}
	fmt.Printf("configuration Loaded")
	return
}
