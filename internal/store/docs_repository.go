package store

import "docs/internal/model"

//DocsRepository - интерфейс, который реализует методы для работы с БД
type DocsRepository interface {
	InsertDoc(*model.Document) error
	GetDocs() (*[]model.Document, error)
	SearchDoc(int) *model.Document
}
