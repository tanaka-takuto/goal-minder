package config

import (
	"github.com/caarlos0/env"
)

// config 設定
type config struct {
	Port      int    `env:"PORT" envDefault:"3000"`
	Env       string `env:"ENV,required"`
	SecretKey string `env:"SECRET_KEY,required"`
}

// cfg 設定
var cfg config

func init() {
	err := env.Parse(&cfg)
	if err != nil {
		panic(err)
	}
}

// Port ポート番号
func Port() int {
	return cfg.Port
}

// SecretKey 秘密鍵
func SecretKey() string {
	return cfg.SecretKey
}
