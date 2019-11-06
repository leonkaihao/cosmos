package redis

import "github.com/leonkaihao/cosmos/pkg/config"

// Config redigo config
type Config struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

// GetConfig gets redigo config
func GetConfig(v *config.Config) (c Config, err error) {
	err = v.UnmarshalKey("redis", &c)
	return
}
