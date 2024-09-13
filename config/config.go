package config

import "github.com/spf13/viper"

var Config *Configuration

type Configuration struct {
	Name        string      `mapstructure:"name"`
	Host        string      `mapstructure:"host"`
	Port        int         `mapstructure:"port"`
	MySQLConfig MySQLConfig `mapstructure:"mysql"`
}

type MySQLConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}

type RedisConfig struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

func init() {
	configFile := "config/config.yaml"

	v := viper.New()
	v.SetConfigFile(configFile)

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	Config = &Configuration{}
	if err := v.Unmarshal(Config); err != nil {
		panic(err)
	}
}
