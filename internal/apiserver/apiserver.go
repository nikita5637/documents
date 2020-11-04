package apiserver

import (
	"database/sql"
	"docs/internal/store/postgresql"
	"fmt"
	"net/http"
	"reflect"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/sirupsen/logrus"
)

//Start запускает сервер
func Start(config *Config, logger *logrus.Logger) error {
	c := reflect.Indirect(reflect.ValueOf(config))
	logger.Debugln("Starting apiserver with next config values:")
	for i := 0; i < c.NumField(); i++ {
		logger.Debugf("\t%s = %v\n", c.Type().Field(i).Name, c.Field(i).Interface())
	}

	db, err := newDB("postgres", config, logger)
	if err != nil {
		return err
	}

	store := postgresql.New(db)

	server := newServer(config, logger, store)

	return http.ListenAndServe(config.BindAddr, server)
}

func newDB(driverName string, config *Config, logger *logrus.Logger) (*sql.DB, error) {
	dataSourceName := fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable", config.DBUser, config.DBPassword, config.DBAddr, config.DBPort, config.DBName)

	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	if err := db.Ping(); err != nil {
		return nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, err
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", config.MigrationsDir), config.DBName, driver)
	if err != nil {
		return nil, err
	}

	logger.Info("Applying migrations")
	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return nil, err
		}
		logger.Info("No migrations changes")
	}

	return db, nil
}
