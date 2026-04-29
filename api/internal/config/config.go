package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	DBHost        string `env:"DBHOST,required"`
	DBUser        string `env:"DBUSER,required"`
	DBPassword    string `env:"DBPASSWORD,required"`
	DBName        string `env:"DBNAME,required"`
	DBPort        string `env:"DBPORT,required"`
	DBSSLMode     string `env:"DBSSLMode,required"`
	ServerPort    string `env:"SERVERPORT,required"`
	RedisAddr     string `env:"REDISADDR,required"`
	RedisPassword string `env:"REDISPASSWORD,required"`
	RedisDB       string `env:"REDISDB,required"`
}

func LoadConfig() *Config {
	var cfg Config
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf("Unable to load config: %v", err)
	}
	return &cfg
}
