package util

import (
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Config is models that represent configuration on application
type Config struct {
	ServerPort             string
	DatabasePort           string
	DatabaseUser           string
	DatabasePassword       string
	DatabaseHost           string
	DatabaseName           string
}

var config Config

//GetConfig : Get Current Config
func GetConfig() Config {
	return config
}

//SetConfig set configuraton
func SetConfig(cfg Config) {
	config = cfg
}

//OnConfigurationChange : Set on configuration file change
func (config *Config) OnConfigurationChange(onConfigChange func(e fsnotify.Event)) {
	viper.OnConfigChange(onConfigChange)
}

func LoadConfig() *Config {
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	viper.SetConfigName("config.yml")

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		logrus.Errorf("Error reading config file %v", err)
	}

	err := viper.Unmarshal(&cfg)
	if err != nil {
		logrus.Errorf("Unable to decode into struct, %v", err)
	}
	viper.WatchConfig()
	SetConfig(cfg)
	return &cfg
}