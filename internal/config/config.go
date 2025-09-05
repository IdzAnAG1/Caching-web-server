package config

import (
	"CachingWebServer/internal/app/cli"
	"bytes"
	_ "embed"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
	"time"
)

type Config struct {
	Server struct {
		Port     int           `mapstructure:"Port" validate:"required,min=0,max=65535"`
		Token    string        `mapstructure:"Token" validate:"required"`
		TokenTTL time.Duration `mapstructure:"TokenTTL" validate:"required"`
	} `mapstructure:"Server"`
	Admin struct {
		Token string `mapstructure:"Token" validate:"required"`
	} `mapstructure:"Admin"`
	
	Database struct {
		URL     string        `mapstructure:"URL" validate:"required"`
		ConnTTL time.Duration `mapstructure:"ConnectionTTL" validate:"required"`
	} `mapstructure:"Database"`

	Logger struct {
		Level string `mapstructure:"Level" validate:"required,oneof=debug local info"`
	} `mapstructure:"Logger"`
}

//go:embed Defaults.yaml
var Defaults []byte

funac MustLoadConfig() Config {
	cfg, err := LoadConfig()
	if err != nil {
		panic(err)
	}
	return cfg
}
func LoadConfig() (Config, error) {
	validate := validator.New()
	var (
		cfg = Config{}
	)
	if err := LoadDefaults(); err != nil {
		return Config{}, err
	}
	if *cli.CfgPath != "" {
		viper.SetConfigFile(*cli.CfgPath)
		if err := viper.MergeInConfig(); err != nil {
			return Config{}, err
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix("CWS")

	if err := viper.Unmarshal(&cfg); err != nil {
		return Config{}, err
	}
	if err := validate.Struct(cfg); err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func LoadDefaults() error {
	viper.SetConfigType("yaml")
	return viper.ReadConfig(bytes.NewReader(Defaults))
}
