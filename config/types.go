package config

import (
	"github.com/spf13/viper"
)

type IConfig interface {
	ParseFromViper(v *viper.Viper) error
}

type IConfigProvider interface {
	Parse(config IConfig) error
}

type IServiceConfig interface {
	Start(cp IConfigProvider) error
}
