package teststore

import (
	"docs/internal/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocsRepository_GetDocs(t *testing.T) {
	doc := model.TestDocument()
	store := New()

	err := store.Docs().InsertDoc(doc)
	assert.Nil(t, err)

	err = store.Docs().InsertDoc(doc)
	assert.NotNil(t, err)

	doc = model.TestDocument()
	doc.Number = 17
	err = store.Docs().InsertDoc(doc)
	assert.Nil(t, err)

	doc = model.TestDocument()
	doc.Number = 3
	err = store.Docs().InsertDoc(doc)
	assert.Nil(t, err)

	doc = model.TestDocument()
	doc.Number = 4
	err = store.Docs().InsertDoc(doc)
	assert.NotNil(t, err)

	docs, err := store.Docs().GetDocs()
	assert.Nil(t, err)
	assert.Equal(t, 3, len(*(docs)))
}
