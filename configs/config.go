package configs

import (
	"github.com/spf13/viper"
)

type conf struct {
	MongoDbUri    string `mapstructure:"MONGODB_URI"`
	Environment   string `mapstructure:"ENVIRONMENT"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
}

type Config interface {
	GetMongoDbUri() string
	GetEnvironment() string
	GetLogLevel() string
	GetWebServerPort() string
}

func LoadConfig(filePath string) (Config, error) {
	conf, err := loadConfig(filePath)
	if err != nil {
		return nil, err
	}
	return conf, nil
}

func loadConfig(path string) (Config, error) {
	var cfg *conf

	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigName(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *conf) GetMongoDbUri() string {
	return c.MongoDbUri
}

func (c *conf) GetEnvironment() string {
	return c.Environment
}

func (c *conf) GetLogLevel() string {
	return c.LogLevel
}

func (c *conf) GetWebServerPort() string {
	return c.WebServerPort
}
