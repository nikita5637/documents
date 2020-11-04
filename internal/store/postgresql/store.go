package postgresql

import (
	"database/sql"
	"docs/internal/store"
)

//Store ...
type Store struct {
	db             *sql.DB
	docsRepository *DocsRepository
}

//New ...
func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

//Docs ...
func (s *Store) Docs() store.DocsRepository {
	if s.docsRepository != nil {
		return s.docsRepository
	}

	s.docsRepository = &DocsRepository{
		store: s,
	}

	return s.docsRepository
}
