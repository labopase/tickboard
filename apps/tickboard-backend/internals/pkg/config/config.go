package config

type Config struct {
	Name        string `mapstructure:"name"`
	Version     string `mapstructure:"version"`
	Debug       bool   `mapstructure:"debug"`
	Environment string `mapstructure:"environment"`
}
