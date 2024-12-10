package gorm

import (
	"github.com/lombard-finance/go-common/config"
	"github.com/spf13/viper"
)

type Config struct {
	Address        string
	User           string
	Password       string
	Database       string
	Port           uint32
	MigrationsPath string
}

func (c *Config) ParseFromViper(v *viper.Viper) error {
	c.User = config.ViperGetOrDefault(v, "database.user", "postgres")
	c.Password = config.ViperGetOrDefault(v, "database.password", "12345")
	c.Database = config.ViperGetOrDefault(v, "database.database", "apibackend")
	c.Port = config.ViperGetOrDefaultUint32(v, "database.port", 5432)
	c.MigrationsPath = config.ViperGetOrDefault(v, "database.migrations_path", "file://migrations")
	c.Address = config.ViperGetOrDefault(v, "database.address", "localhost")
	return nil
}
