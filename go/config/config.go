package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	PgHost     string `mapstructure:"PG_HOST"`
	PgPort     int    `mapstructure:"PG_PORT"`
	PgDBName   string `mapstructure:"PG_DB"`
	PgUser     string `mapstructure:"PG_USER"`
	PgPassword string `mapstructure:"PG_PASSWORD"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	JwtSecret  string `mapstructure:"JWT_SECRET"`
	SslMode    string `mapstructure:"SSL_MODE"`
	PGDump     string
	Schemas    []string
	NoHeader   bool
	Writer     *os.File
	NoClean    bool
}

var Cfg Config

func LoadConfig() (*Config, error) {
	viper.AddConfigPath("./")
	if os.Getenv("ENV") == "prod" {
		viper.SetConfigName(".env.prod")
	} else if os.Getenv("ENV") == "staging" {
		viper.SetConfigName(".env.staging")
	} else {
		viper.SetConfigName(".env.dev")
	}
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	viper.AutomaticEnv()
	if err != nil {
		return &Cfg, err
	}
	err = viper.Unmarshal(&Cfg)
	return &Cfg, err
}
