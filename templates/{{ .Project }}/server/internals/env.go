package internals

import (
	"fmt"

	"github.com/samber/do"
	"github.com/spf13/viper"
)

type Env struct {
	DATABASE_URL string
}

func NewEnv(i *do.Injector) (*Env, error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		return nil, fmt.Errorf("fatal error config file: %w", err)
	}

	return &Env{
		DATABASE_URL: viper.GetString("DATABASE_URL"),
	}, nil
}
