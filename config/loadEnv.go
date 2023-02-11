package config

import (
	"log"

	"github.com/spf13/viper"
)

type ConfigDB1 struct {
	DBHost1         string `mapstructure:"POSTGRES_HOST1"`
	DBUserName1     string `mapstructure:"POSTGRES_USER1"`
	DBUserPassword1 string `mapstructure:"POSTGRES_PASSWORD1"`
	DBName1         string `mapstructure:"POSTGRES_DB1"`
	DBPort1         string `mapstructure:"POSTGRES_PORT1"`
	ServerPort1     string `mapstructure:"PORT1"`
}

type ConfigDB2 struct {
	DBHost2         string `mapstructure:"POSTGRES_HOST2"`
	DBUserName2     string `mapstructure:"POSTGRES_USER2"`
	DBUserPassword2 string `mapstructure:"POSTGRES_PASSWORD2"`
	DBName2         string `mapstructure:"POSTGRES_DB2"`
	DBPort2         string `mapstructure:"POSTGRES_PORT2"`
	ServerPort2     string `mapstructure:"PORT2"`
}

var DB1 ConfigDB1
var DB2 ConfigDB2

func init() {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AddConfigPath(".")

	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		log.Panic(err)
	}

	if err := v.Unmarshal(&DB1); err != nil {
		log.Panic(err)
	}

	if err := v.Unmarshal(&DB2); err != nil {
		log.Panic(err)
	}
}
