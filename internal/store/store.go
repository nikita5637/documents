package store

//Store - интерфейс, который реализует методы, возвращающие интерфейсы для работы с БД
type Store interface {
	Docs() DocsRepository
}
