package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// global config
var Cfg Config

type (
	Postgres struct {
		DB   string `mapstructure:"db"   validate:"required"`
		User string `mapstructure:"user" validate:"required"`
		Pass string `mapstructure:"pass" validate:"required"`
		Host string `mapstructure:"host" validate:"required"`
		Port string `mapstructure:"port" validate:"required"`
	}

	Config struct {
		Port     int      `mapstructure:"port"`
		Postgres Postgres `mapstructure:"postgres"`
	}
)

func Init(path string) {
	v := viper.New()
	v.SetConfigType("yml")
	v.SetConfigName("config")
	v.AddConfigPath(path)
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Fatal error config file: %s \n", err)
	}

	if err := v.UnmarshalExact(&Cfg); err != nil {
		logrus.Fatalf("unmarshaling error: %s", err)
	}

	logrus.Printf("config loaded: %+v", Cfg)
}
