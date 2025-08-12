package config

import (
	"CachingWebServer/internal/flags"
	"bytes"
	_ "embed"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"Port" validate:"required,min=0,max=65535"`
	} `mapstructure:"Server"`

	Database struct {
		URL string `mapstructure:"URL"`
	} `mapstructure:"Database"`

	Logger struct {
		Level string `mapstructure:"Level"`
	} `mapstructure:"Logger"`
}

//go:embed Defaults.yaml
var Defaults []byte

func MustLoadConfig() Config {
	validate := validator.New()
	var (
		cfg = Config{}
	)
	if err := LoadDefaults(); err != nil {
		panic("Cannot Load Defaults fields")
	}
	if *flags.CfgPath != "" {
		viper.SetConfigFile(*flags.CfgPath)
	}
	viper.AutomaticEnv()
	viper.SetEnvPrefix("CWS")

	if err := viper.Unmarshal(&cfg); err != nil {
		panic("")
	}
	if err := validate.Struct(cfg); err != nil {
		panic("")
	}
	return cfg
}

func LoadDefaults() error {
	viper.SetConfigType("yaml")
	if err := viper.ReadConfig(bytes.NewReader(Defaults)); err != nil {
		return errors.New(fmt.Sprintf("Failed to read configuration defaults fields (%v)", err))
	}
	return nil
}
