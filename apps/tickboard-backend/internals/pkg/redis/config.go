package redis

import (
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	poolSize     = 10
	minIdleConns = 5
	dialTimeout  = 5 * time.Second
	readTimeout  = 3 * time.Second
	writeTimeout = 3 * time.Second
)

type Config struct {
	Addrs        []string      `mapstructure:"addrs"`
	Password     string        `mapstructure:"password"`
	DB           int           `mapstructure:"db"`
	PoolSize     int           `mapstructure:"pool_size"`
	MinIdleConns int           `mapstructure:"min_idle_conns"`
	DialTimeout  time.Duration `mapstructure:"dial_timeout"`
	ReadTimeout  time.Duration `mapstructure:"read_timeout"`
	WriteTimeout time.Duration `mapstructure:"write_timeout"`
}

func (c *Config) redisOption() *redis.UniversalOptions {
	opt := &redis.UniversalOptions{
		Addrs:        c.Addrs,
		Password:     c.Password,
		PoolSize:     c.PoolSize,
		MinIdleConns: c.MinIdleConns,
		DialTimeout:  c.DialTimeout,
		ReadTimeout:  c.ReadTimeout,
		WriteTimeout: c.WriteTimeout,
	}

	if len(c.Addrs) == 1 {
		opt.DB = c.DB
	}

	return opt
}

func (c *Config) applyDefaults() {
	if len(c.Addrs) == 0 {
		c.Addrs = []string{"127.0.0.1:6379"}
	}

	if c.PoolSize == 0 {
		c.PoolSize = poolSize
	}

	if c.MinIdleConns == 0 {
		c.MinIdleConns = minIdleConns
	}

	if c.DialTimeout == 0 {
		c.DialTimeout = dialTimeout
	}

	if c.ReadTimeout == 0 {
		c.ReadTimeout = readTimeout
	}

	if c.WriteTimeout == 0 {
		c.WriteTimeout = writeTimeout
	}
}
