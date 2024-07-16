package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type DBConfig struct {
	Host        string `yaml:"host"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Database    string `yaml:"dbname"`
	Port        int    `yaml:"port"`
	AutoMigrate bool   `yaml:"autoMigrate"`
	Timezone    string `yaml:"timezone"`
	LogLevel    string `yaml:"logLevel"`
	// SlowThreshold marshaltype.Duration `yaml:"slowThreshold"`
}
