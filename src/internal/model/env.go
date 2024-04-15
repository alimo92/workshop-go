package model

import (
	"log"

	"github.com/spf13/viper"
)

var Env Environment

type Environment struct {
	// Development environment. e.g. local/prod
	DevEnv  string `mapstructure:"DEV_ENV"`
	AppPort int    `mapstructure:"APP_PORT"`

	// Postgres
	PostgresHost     string `mapstructure:"POSTGRES_HOST"`
	PostgresUser     string `mapstructure:"POSTGRES_USER"`
	PostgresPassword string `mapstructure:"POSTGRES_PASSWORD"`
	PostgresDb       string `mapstructure:"POSTGRES_DB"`
	PostgresSslMode  string `mapstructure:"POSTGRES_SSL_MODE"`
	PostgresPort     int    `mapstructure:"POSTGRES_PORT"`

	// Slack
	SlackBotToken string `mapstructure:"SLACK_BOT_TOKEN"`
	SlackChannel  string `mapstructure:"SLACK_CHANNEL"`
	SlackAppToken string `mapstructure:"SLACK_APP_TOKEN"`
}

func init() {
	viper.SetConfigName(".env")
	viper.AddConfigPath("./../../../.")
	viper.AddConfigPath("./")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't read environment variables : ", err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
}
