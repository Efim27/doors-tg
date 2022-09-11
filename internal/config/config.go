package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	TokenTG        string `mapstructure:"TG_TOKEN"`
	IsWebhook      bool   `mapstructure:"IS_WEBHOOK"`
	AppPort        string `mapstructure:"APP_PORT"`
	AppAddr        string `mapstructure:"APP_ADDR"`
	ControlAPIAddr string `mapstructure:"CONTROL_API_ADDR"`
	Debug          bool   `mapstructure:"DEBUG"`
}

func LoadConfig(path string) (config Config, err error) {
	// Config default values
	viper.SetDefault("TG_TOKEN", "")
	viper.SetDefault("IS_WEBHOOK", false)
	viper.SetDefault("APP_PORT", "")
	viper.SetDefault("APP_ADDR", "")
	viper.SetDefault("CONTROL_API_ADDR", "")
	viper.SetDefault("DEBUG", false)

	viper.AddConfigPath(path)
	viper.SetConfigName("main")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
