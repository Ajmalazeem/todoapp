package util

import "github.com/spf13/viper"

type Config struct {
	DBHost string `mapstructure:"DB_HOST"`
	DBName string `mapstructure:"DB_NAME"`
	DBUser string `mapstructure:"DB_USER"`
	DBPort int `mapstructure:"DB_PORT"`
	DBPassword int `mapstructure:"DB_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil{
		return
	}

	err = viper.Unmarshal(&config)
	return
}