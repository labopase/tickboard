package logger

type Config struct {
	Level        string `mapstructure:"level"`
	Environment  string `mapstructure:"environment"`
	EnableCaller bool   `mapstructure:"enable_caller"`
	EnableTrace  bool   `mapstructure:"enable_trace"`
}
