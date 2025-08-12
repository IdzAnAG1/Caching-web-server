package config

type Config struct {
	Server struct {
		Port int `mapstructure:"Port"`
	}
}

func LoadConfig() (Config, error) {

}
