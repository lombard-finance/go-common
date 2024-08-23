package logger

import (
	"github.com/lombard-finance/go-common/config"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	// trace, debug, info, warning, error, fatal, panic
	LogLevel         logrus.Level
	DisableTimestamp bool
	JsonFormatter    bool
	PrettyPrint      bool
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.DisableTimestamp = config.ViperGetOrDefaultBool(v, "logger.disable-timestamp", false)
	c.JsonFormatter = config.ViperGetOrDefaultBool(v, "logger.json-formatter", false)
	c.PrettyPrint = config.ViperGetOrDefaultBool(v, "logger.pretty-print", false)

	logLevel := config.ViperGetOrDefault(v, "logger.log-level", logrus.InfoLevel.String())

	var err error
	c.LogLevel, err = logrus.ParseLevel(logLevel)
	if err != nil {
		return errors.Wrap(err, "invalid logger.log-level")
	}
	return nil
}
