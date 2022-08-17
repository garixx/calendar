package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"strings"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func GetPostgresConfig(path string) (PostgresConfig, error) {
	err := initConfig(path)
	if err != nil {
		return PostgresConfig{}, err
	}
	return PostgresConfig{
		Host:     viper.GetString("database.postgres.host"),
		Port:     viper.GetString("database.postgres.port"),
		Username: viper.GetString("database.postgres.user"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   viper.GetString("database.postgres.dbname"),
		SSLMode:  viper.GetString("database.postgres.sslmode"),
	}, nil
}

func initConfig(path string) error {
	//// make qa environment default
	//_ = viper.BindEnv("environment")
	//viper.SetDefault("environment", "QA")

	path = fmt.Sprintf(path, strings.ToLower(os.Getenv("environment")))
	viper.SetConfigFile(path)
	return viper.ReadInConfig()
}
