package config

type Config struct {
	Server ServerConfig `yaml:"server"`
	DB     DBConfig     `yaml:"db"`
	Log    LogConfig    `yaml:"log"`
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

type LogConfig struct {
	Mode       string `yaml:"mode"`
	Format     string `yaml:"format"`
	File       string `yaml:"file"`
	Level      string `yaml:"level"`
	MaxSize    int    `yaml:"maxSize"`
	MaxAge     int    `yaml:"maxAge"`
	MaxBackups int    `yaml:"maxBackups"`
}
