package config

import (
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	App        App
	HttpServer HttpServer
	Chat       Chat
}

type App struct {
	Name string
}

type HttpServer struct {
	RunModel     string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Chat struct {
	Token string
}

var Conf Config

func Setup() {

	config := new(Config)

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatal("read yaml file err:", err)
		}
	}

	if err := viper.Unmarshal(&config); err != nil {
		if _, ok := err.(viper.ConfigMarshalError); ok {
			log.Fatal("unmarshal yaml err:", err)
		}
	}

	Conf = *config

}
