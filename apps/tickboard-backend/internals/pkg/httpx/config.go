package httpx

import (
	"fmt"
	"time"
)

const (
	writeTimeout    = 15 * time.Second
	readTimeout     = 15 * time.Second
	idleTimeout     = 30 * time.Second
	shutdownTimeout = 30 * time.Second
	maxHeaderBytes  = 1 << 20 // 1MB
	maxBodyBytes    = 1 << 20
)

type Config struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`

	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
	MaxHeaderBytes  int           `mapstructure:"max_header_bytes"`
	MaxBodyBytes    int           `mapstructure:"max_body_bytes"`
}

func (c *Config) Addr() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *Config) applyDefaults() {
	if c.WriteTimeout == 0 {
		c.WriteTimeout = writeTimeout
	}
	if c.ReadTimeout == 0 {
		c.ReadTimeout = readTimeout
	}
	if c.IdleTimeout == 0 {
		c.IdleTimeout = idleTimeout
	}
	if c.ShutdownTimeout == 0 {
		c.ShutdownTimeout = shutdownTimeout
	}
	if c.MaxHeaderBytes == 0 {
		c.MaxHeaderBytes = maxHeaderBytes
	}
	if c.MaxBodyBytes == 0 {
		c.MaxBodyBytes = maxBodyBytes
	}
}
