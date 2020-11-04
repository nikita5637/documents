package apiserver

import (
	"log"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	if err := os.Setenv("Testing", "true"); err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}
