package postgresql

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

//TestDB возвращает подключение к тестовой базе и функцию по очистке таблиц после выполнения тестов
func TestDB(t *testing.T, dataSourceName string) (*sql.DB, func(...string), error) {
	t.Helper()

	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, nil, err
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return nil, nil, err
	}

	var migrationsDir string
	if migrationsDir = os.Getenv("POSTGRESQL_MIGRATIONS_DIR"); migrationsDir == "" {
		migrationsDir = "/go/src/docs/internal/migrations/"
	}

	m, err := migrate.NewWithDatabaseInstance(fmt.Sprintf("file://%s", migrationsDir), "docs_test", driver)
	if err != nil {
		return nil, nil, err
	}

	if err := m.Up(); err != nil {
		if err != migrate.ErrNoChange {
			return nil, nil, err
		}
	}

	return db, func(tables ...string) {
		if len(tables) > 0 {
			for _, table := range tables {
				db.Exec(fmt.Sprintf("TRUNCATE TABLE %s", table))
			}
		}
	}, nil
}
