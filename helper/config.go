package helper

import "github.com/spf13/viper"

type Config struct {
	DB           string `mapstructure:"DB"`
	SCHEMA       string `mapstructure:"SCHEMA"`
	GIN_MODE     string `mapstructure:"GIN_MODE"`
	ENV          string `mapstructure:"ENV"`
	LOG_FILE     string `mapstructure:"LOG_FILE"`
	AUTO_MIGRATE string `mapstructure:"AUTO_MIGRATE"`
	PORT         string `mapstructure:"PORT"`
	ALLOW_ORIGIN string `mapstructure:"ALLOW_ORIGIN"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()

	if err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}