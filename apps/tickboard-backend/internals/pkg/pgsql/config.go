package pgsql

import (
	"fmt"
	"time"
)

type Config struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	User            string        `mapstructure:"user"`
	Password        string        `mapstructure:"password"`
	DBName          string        `mapstructure:"dbname"`
	SSLMode         string        `mapstructure:"ssl_mode"`
	ConnectTimeout  time.Duration `mapstructure:"connect_timeout"`
	ApplicationName string        `mapstructure:"application_name"`

	MaxOpenConns    int32         `mapstructure:"max_open_conns"`
	MaxIdleConns    int32         `mapstructure:"max_idle_conns"`
	ConnMaxLifetime time.Duration `mapstructure:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `mapstructure:"conn_max_idle_time"`
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=%s&connect_timeout=%.0f&application_name=%s", c.User, c.Password, c.Host, c.Port, c.DBName, c.SSLMode, c.ConnectTimeout.Seconds(), c.ApplicationName)
}

func (c *Config) applyDefaults() {
	if c.ConnectTimeout == 0 {
		c.ConnectTimeout = 30 * time.Second
	}
	if c.MaxOpenConns == 0 {
		c.MaxOpenConns = 10
	}
	if c.MaxIdleConns == 0 {
		c.MaxIdleConns = 5
	}
	if c.ConnMaxLifetime == 0 {
		c.ConnMaxLifetime = 1 * time.Hour
	}
	if c.ConnMaxIdleTime == 0 {
		c.ConnMaxIdleTime = 1 * time.Hour
	}
}
