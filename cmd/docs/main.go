package main

import (
	"docs/internal/apiserver"
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config", "./configs/docs.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	if _, err := toml.DecodeFile(configPath, &config); err != nil {
		log.Fatal(err)
	}

	for _, env := range os.Environ() {
		switch strings.Split(env, "=")[0] {
		case "POSTGRESQL_IP_ADDRESS":
			config.DBAddr = os.Getenv("POSTGRESQL_IP_ADDRESS")
		case "POSTGRESQL_PORT":
			str := os.Getenv("POSTGRESQL_PORT")
			if i, err := strconv.Atoi(str); err == nil {
				config.DBPort = uint16(i)
			}
		case "POSTGRESQL_USER":
			config.DBUser = os.Getenv("POSTGRESQL_USER")
		case "POSTGRESQL_PASSWORD":
			config.DBPassword = os.Getenv("POSTGRESQL_PASSWORD")
		case "POSTGRESQL_DBNAME":
			config.DBName = os.Getenv("POSTGRESQL_DBNAME")
		case "POSTGRESQL_MIGRATIONS_DIR":
			config.MigrationsDir = os.Getenv("POSTGRESQL_MIGRATIONS_DIR")
		case "DOCS_BIND_ADDR":
			config.BindAddr = os.Getenv("DOCS_BIND_ADDR")
		case "DOCS_LOG_LEVEL":
			config.LogLevel = os.Getenv("DOCS_LOG_LEVEL")
		}
	}

	if err := config.Validate(); err != nil {
		log.Fatal(err)
	}

	logger := logrus.New()
	level, _ := logrus.ParseLevel(config.LogLevel)
	logger.SetLevel(level)

	if err := apiserver.Start(config, logger); err != nil {
		logger.Fatal(err)
	}
}
