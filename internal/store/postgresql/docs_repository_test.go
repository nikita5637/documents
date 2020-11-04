package postgresql_test

import (
	"docs/internal/model"
	"docs/internal/store/postgresql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDocsRepository_TestDatabaseQueries(t *testing.T) {
	db, teardown, err := postgresql.TestDB(t, dataSourceName)
	assert.Nil(t, err)

	defer teardown("documents")

	store := postgresql.New(db)
	doc := model.TestDocument()

	err = store.Docs().InsertDoc(doc)
	assert.Nil(t, err)

	//Duplicate number
	err = store.Docs().InsertDoc(doc)
	assert.NotNil(t, err)

	//Insert document with number 2
	doc = model.TestDocument()
	doc.Number = 2
	err = store.Docs().InsertDoc(doc)
	assert.Nil(t, err)

	//Get all documents
	docs, err := store.Docs().GetDocs()
	assert.NotNil(t, docs)
	assert.Nil(t, err)

	//Search existing document
	getDoc := store.Docs().SearchDoc(2)
	assert.NotNil(t, getDoc)

	//Search not existing document
	getDoc = store.Docs().SearchDoc(3)
	assert.Nil(t, getDoc)
}
