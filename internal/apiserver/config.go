package apiserver

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/sirupsen/logrus"
)

//Config структура, хранящая переменные конфигурации
// BindAddr - локальный адрес для получения запросов
// LogLevel - уровень логирования
// MinDocsPerRequest - минимальное количество документов в запросе
// MaxDocsPerRequest int64 - максимальное количество документов в запросе
// DBAddr - адрес БД
// DBUser - имя пользователя БД
// DBPassword - пароль пользователя БД
// DBName - имя базы БД
// MigrationsDir - папка с миграциями
type Config struct {
	BindAddr          string `toml:"bind_addr"`
	LogLevel          string `toml:"log_level"`
	MinDocsPerRequest int64  `toml:"min_docs_per_request"`
	MaxDocsPerRequest int64  `toml:"max_docs_per_request"`
	DBAddr            string `toml:"db_addr"`
	DBPort            uint16 `toml:"db_port"`
	DBUser            string `toml:"db_user"`
	DBPassword        string `toml:"db_password"`
	DBName            string `toml:"db_name"`
	MigrationsDir     string `toml:"migrations_dir"`
}

//NewConfig возвращает указатель на структуру конфига, с заполненными дефолтными значениями
func NewConfig() *Config {
	return &Config{
		BindAddr:          ":80",
		LogLevel:          "info",
		MinDocsPerRequest: 1,
		MaxDocsPerRequest: 10,
	}
}

//Validate валидирует конфиг сервиса
func (c *Config) Validate() error {
	if err := validation.ValidateStruct(c,
		validation.Field(&c.DBAddr, validation.Required, is.IPv4),
		validation.Field(&c.DBPort, validation.Required, validation.Min(uint16(1)), validation.Max(uint16(65535))),
		validation.Field(&c.DBUser, validation.Required, is.Alphanumeric),
		validation.Field(&c.DBPassword, validation.Required, is.PrintableASCII),
		validation.Field(&c.DBName, validation.Required, is.Alphanumeric),
		validation.Field(&c.MigrationsDir, validation.Required),

		validation.Field(&c.MinDocsPerRequest, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&c.MaxDocsPerRequest, validation.Required, validation.Min(1), validation.Max(100)),
		validation.Field(&c.BindAddr, validation.Required),
		validation.Field(&c.LogLevel, validation.Required, is.Alpha),
	); err != nil {
		return err
	}

	if c.MinDocsPerRequest > c.MaxDocsPerRequest {
		return errors.New("mid_docs_per_request must be less than max_docs_per_request")
	}

	if _, err := logrus.ParseLevel(c.LogLevel); err != nil {
		return err
	}

	return nil
}
