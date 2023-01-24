package config

import (
	"fmt"
	"github.com/spf13/viper"
	"net/url"
)

type DatabaseConfig struct {
	DBHost     string `mapstructure:"POSTGRES_HOST"`
	DBUser     string `mapstructure:"POSTGRES_USER"`
	DBPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBPort     uint16 `mapstructure:"POSTGRES_PORT"`
	DBName     string `mapstructure:"POSTGRES_DB"`
}

func LoadDatabaseConfig(path string) (config DatabaseConfig, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("db")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err = viper.ReadInConfig(); err != nil {
		return
	}
	err = viper.Unmarshal(&config)
	return
}

func (dbConfig *DatabaseConfig) ToConnectionString() string {
	dsn := url.URL{
		User:     url.UserPassword(dbConfig.DBUser, dbConfig.DBPassword),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", dbConfig.DBHost, dbConfig.DBPort),
		Path:     dbConfig.DBName,
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	return dsn.String()
}
