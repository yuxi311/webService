package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

var (
	internal_config Config
)

func Load(configFile string) error {
	if _, err := os.Stat(configFile); err != nil {
		return err
	}

	cfg := Config{}

	// read config file webserver.yaml
	bytes, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(bytes, &cfg)
	if err != nil {
		return err
	}

	internal_config = cfg
	return nil
}

func Get() Config {
	return internal_config
}

func Server() ServerConfig {
	return internal_config.Server
}

func DB() DBConfig {
	return internal_config.DB
}

func Log() LogConfig {
	return internal_config.Log
}

func Redis() RedisConfig {
	return internal_config.Redis
}

func Kafka() KafkaConfig {
	return internal_config.Kafka
}

func MQTT() MQTTConfig {
	return internal_config.MQTT
}
