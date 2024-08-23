package logger

import (
	"github.com/lombard-finance/go-common/config"
	"github.com/sirupsen/logrus"
	"os"
)

func NewLogger(cp config.IConfigProvider) *logrus.Logger {
	log := logrus.New()
	// Set the output to stdout (default is stderr)
	log.Out = os.Stdout

	cfg := &Config{}
	if err := cp.Parse(cfg); err != nil {
		log.Fatalf("failed to parse config: %v", err)
	}
	log.SetLevel(cfg.LogLevel)

	var formatter logrus.Formatter = &logrus.TextFormatter{
		ForceColors:      true,
		DisableTimestamp: cfg.DisableTimestamp,
		FullTimestamp:    true,
		PadLevelText:     true,
		QuoteEmptyFields: true,
	}

	if cfg.JsonFormatter {
		formatter = &logrus.JSONFormatter{
			DisableTimestamp: cfg.DisableTimestamp,
			PrettyPrint:      cfg.PrettyPrint,
		}
	}

	log.SetFormatter(formatter)

	return log
}
