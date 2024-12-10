package gorm

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDSN(addr, user, pass, db string, port uint32) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=UTC",
		addr, user, pass, db, port,
	)
}

func runMigrations(config Config, log *logrus.Entry) {
	m, err := migrate.New(
		config.MigrationsPath,
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Address, config.Port, config.Database))
	if err != nil {
		log.WithError(err).Fatalf("can't init migration")
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			log.Info("no changes applied to the migration")
		} else {
			log.WithError(err).Fatalf("unable to apply migrations")
		}
	}
	version, dirty, err := m.Version()
	if err != nil {
		log.WithError(err).
			WithField("version", version).
			WithField("dirty", dirty).
			Fatalf("get migration version")
	}
	log.Infof("current migration version: %d\n", version)

	sErr, dbErr := m.Close()
	if dbErr != nil {
		log.WithError(dbErr).Fatal("can't close migration")
	}
	if sErr != nil {
		log.WithError(dbErr).Fatal("can't close migration")
	}
}

func NewDriver(config Config, log *logrus.Entry) (*gorm.DB, error) {
	runMigrations(config, log)

	db, err := gorm.Open(
		postgres.Open(NewDSN(config.Address, config.User, config.Password, config.Database, config.Port)),
		&gorm.Config{
			Logger:         logger.Default.LogMode(logger.Silent),
			TranslateError: true,
		},
	)
	if err != nil {
		return nil, errors.Wrap(err, "open connection")
	}
	return db, nil
}
