package postgresql_test

import (
	"docs/internal/store/postgresql"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dataSourceName string
)

func TestMain(m *testing.M) {
	DBAddr := "172.20.1.3"
	DBUser := "postgres"
	DBPassword := "password"
	DBName := "docs_test"
	DBPort := uint16(5432)

	for _, env := range os.Environ() {
		switch strings.Split(env, "=")[0] {
		case "POSTGRESQL_IP_ADDRESS":
			DBAddr = os.Getenv("POSTGRESQL_IP_ADDRESS")
		case "POSTGRESQL_PORT":
			str := os.Getenv("POSTGRESQL_PORT")
			if i, err := strconv.Atoi(str); err == nil {
				DBPort = uint16(i)
			}
		case "POSTGRESQL_USER":
			DBUser = os.Getenv("POSTGRESQL_USER")
		case "POSTGRESQL_PASSWORD":
			DBPassword = os.Getenv("POSTGRESQL_PASSWORD")
		case "POSTGRESQL_DBNAME":
			DBName = os.Getenv("POSTGRESQL_DBNAME")
		}
	}

	dataSourceName = fmt.Sprintf("user=%s password=%s host=%s port=%d database=%s sslmode=disable", DBUser, DBPassword, DBAddr, DBPort, DBName)

	os.Exit(m.Run())
}

func TestNew(t *testing.T) {
	store := postgresql.New(nil)
	assert.NotNil(t, store)
}
