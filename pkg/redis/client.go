package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/leonkaihao/cosmos/pkg/log"
)

// type alias
type (
	// Conn ...
	Conn = redis.Conn
	// Pool ...
	Pool = redis.Pool
)

// NewConn creates a new redis connection
func NewConn(c Config) (Conn, error) {
	conn, err := redis.Dial("tcp", c.Address, redis.DialPassword(c.Password))
	if err != nil {
		log.Errorf("Failed to dial %v for redis: %v", c.Address, err)
		return nil, err
	}
	return conn, nil
}

// NewPool creates a new pool
func NewPool(c Config) *Pool {
	return redis.NewPool(
		func() (redis.Conn, error) {
			return NewConn(c)
		}, 3)
}
