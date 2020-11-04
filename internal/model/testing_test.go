package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_TestDocument(t *testing.T) {
	doc := TestDocument()
	assert.NotNil(t, doc)
}
