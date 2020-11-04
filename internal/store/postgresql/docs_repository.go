package postgresql

import (
	"docs/internal/model"
)

//DocsRepository - хранилище
type DocsRepository struct {
	store *Store
}

//InsertDoc - вставка в БД
func (d *DocsRepository) InsertDoc(doc *model.Document) error {
	var lastInsertID int
	if err := d.store.db.QueryRow("INSERT INTO documents (name, date, number, sum) VALUES ($1, $2, $3, $4) RETURNING id",
		doc.Name, doc.Date, doc.Number, doc.Sum).Scan(&lastInsertID); err != nil {
		return err
	}

	return nil
}

//GetDocs - получить все документы из БД
func (d *DocsRepository) GetDocs() (*[]model.Document, error) {
	var docs = make([]model.Document, 0, 1024)
	rows, err := d.store.db.Query("SELECT name, date, number, sum FROM documents")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		doc := model.Document{}
		if err := rows.Scan(&doc.Name, &doc.Date, &doc.Number, &doc.Sum); err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}

	return &docs, nil
}

//SearchDoc - поиск в БД
func (d *DocsRepository) SearchDoc(number int) *model.Document {
	var doc model.Document
	if err := d.store.db.QueryRow("SELECT name, date, number, sum FROM documents WHERE number = $1",
		number).Scan(&doc.Name, &doc.Date, &doc.Number, &doc.Sum); err != nil {
		return nil
	}

	return &doc
}
