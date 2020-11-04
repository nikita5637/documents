package teststore

import (
	"docs/internal/model"
	"fmt"
	"sync"
)

//DocsRepository - тестовое хранилище, реализующее инетрфейс DocsRepository
type DocsRepository struct {
	documents map[int]model.Document
	mutex     sync.Mutex
}

//InsertDoc ...
func (d *DocsRepository) InsertDoc(doc *model.Document) error {
	if doc := d.SearchDoc(doc.Number); doc != nil {
		return fmt.Errorf("Number %d already exists", doc.Number)
	}

	d.mutex.Lock()
	defer d.mutex.Unlock()
	nextID := len(d.documents)
	if nextID >= 3 {
		return fmt.Errorf("No free space")
	}

	d.documents[nextID] = *doc
	return nil
}

//GetDocs ...
func (d *DocsRepository) GetDocs() (*[]model.Document, error) {
	var docs = make([]model.Document, 0, 1024)
	for _, doc := range d.documents {
		docs = append(docs, doc)
	}

	return &docs, nil
}

//SearchDoc ...
func (d *DocsRepository) SearchDoc(number int) *model.Document {
	for _, doc := range d.documents {
		if doc.Number == number {
			return &doc
		}
	}

	return nil
}
