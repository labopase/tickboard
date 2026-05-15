package uconfig

import (
	"fmt"

	"github.com/spf13/viper"
)

const (
	DefaultConfigPath = "./"
	DefaultConfigName = "config"
)

func BindJSONKey[T any](key string) (T, error) {
	return Bind[T](key, DefaultConfigPath, DefaultConfigName, "json")
}

func BindYAMLKey[T any](key string) (T, error) {
	return Bind[T](key, DefaultConfigPath, DefaultConfigName, "yaml")
}

func Bind[T any](key, path, name, extension string) (T, error) {
	if path == "" {
		path = "./"
	}

	if name == "" {
		name = "config"
	}

	if extension == "" {
		extension = "json"
	}

	var result T

	v := viper.New()

	v.SetConfigName(name)
	v.AddConfigPath(path)
	v.SetConfigType(extension)

	if err := v.ReadInConfig(); err != nil {
		return result, fmt.Errorf("viper failed to read config file: %w", err)
	}

	if key == "" {
		if err := v.Unmarshal(&result); err != nil {
			return result, fmt.Errorf("viper failed to unmarshal config file: %w", err)
		}
	} else {
		if err := v.UnmarshalKey(key, &result); err != nil {
			return result, fmt.Errorf("viper unable to decode into struct, %w", err)
		}
	}
	return result, nil
}
