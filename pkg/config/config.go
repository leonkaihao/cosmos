package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config ...
type Config = viper.Viper

// LoadConfig for service based on env, load local data
func LoadConfig(svc string) (v *Config, err error) {
	v = viper.New()
	v.SetConfigFile(fmt.Sprintf("config/%s.yaml", svc))
	err = v.ReadInConfig()
	return
}

// GetLocalConfig gets data from full or related path
func GetLocalConfig(path string) (v *Config, err error) {
	v = viper.New()
	v.SetConfigFile(path)
	err = v.ReadInConfig()
	return
}
