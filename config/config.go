package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DB struct {
		Driver  string `json:"driver"`
		Address string `json:"address"`
	} `json:"db"`
	Address string `json:"address"`
	Mode    string `json:"mode"`
}

var config *Config

func Load() *Config {
	if config != nil {
		return config
	}

	loadConfig()

	return config
}

func loadConfig() {
	viper.SetConfigName("configuration")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln("error reading file config: ", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatalln("error parsing file config: ", err)
	}

	fmt.Println("config.yml", config)
}
