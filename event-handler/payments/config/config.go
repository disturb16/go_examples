package config

type Config struct {
	Address      string `json:"app_address"`
	DBPath       string `json:"db_path"`
	RabbitMQAddr string `json:"rabbitmq_addr"`
	LogLevel     int    `json:"log_level"`
}

func New() *Config {
	// TODO: Read values from config file

	return &Config{
		Address:      ":8082",
		DBPath:       "data.db",
		RabbitMQAddr: "amqp://guest:guest@localhost:5672",
		LogLevel:     -4,
	}
}
