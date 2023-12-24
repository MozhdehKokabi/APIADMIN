package config

import (
	// "os"

	// "fmt"

	"github.com/spf13/viper"
)

var cfg = &config{}

func LoadConfig() error {
	envFile := ".env"
	// if os.Getenv("DEV_MODE") == "1" {
	// 	envFile = ".env"
	// }

	viper.SetConfigFile(envFile)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.Unmarshal(&cfg)
	return err
}

func GetPostgres() *postgres {
	return &cfg.Postgres
}
