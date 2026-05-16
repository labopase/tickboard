package healthz

type Config struct {
	Redis    bool `mapstructure:"redis"`
	Postgres bool `mapstructure:"postgres"`
	Kafka    bool `mapstructure:"kafka"`
	Grpc     bool `mapstructure:"grpc"`
}
