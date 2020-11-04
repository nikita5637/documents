package teststore

import (
	"docs/internal/model"
	"docs/internal/store"
)

//Store - тестовое хранилище документов
type Store struct {
	docsRepository *DocsRepository
}

//New ...
func New() *Store {
	return &Store{}
}

//Docs - возвращает структуру, удовлетворяющую интерфейсу DocsRepository для работы с тестовой БД
func (s *Store) Docs() store.DocsRepository {
	if s.docsRepository != nil {
		return s.docsRepository
	}

	s.docsRepository = &DocsRepository{
		documents: make(map[int]model.Document),
	}

	return s.docsRepository
}
