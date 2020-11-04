package apiserver

import (
	"docs/internal/model"
	"docs/internal/store"
	"fmt"
	"sync"
)

//Список документов в запросе
type requestBody struct {
	Documents []model.Document `json:"documents"`
	store     store.Store
}

//validate валидирует все документы в запросе
//возвращает ошибку, которая появится первой, во время параллельной проверки всех документов в разных горутинах
func (r *requestBody) validate(min, max int64) error {
	numberDocs := int64(len(r.Documents))
	if numberDocs > max || numberDocs < min {
		return fmt.Errorf("The number of documents must be between %d and %d", min, max)
	}

	var wgDoneChan = make(chan bool)
	var errorsChan = make(chan error)
	var wg sync.WaitGroup

	for _, doc := range r.Documents {
		wg.Add(1)
		lDoc := doc
		go func(doc *model.Document) {
			defer wg.Done()
			if err := doc.Validate(); err != nil {
				errorsChan <- err
				return
			}

			if d := r.store.Docs().SearchDoc(doc.Number); d != nil {
				errorsChan <- fmt.Errorf("Number %d already exists", doc.Number)
			}
		}(&lDoc)
	}

	go func() {
		wg.Wait()
		close(wgDoneChan)
	}()

	select {
	case <-wgDoneChan:
		return nil
	case err := <-errorsChan:
		return err
	}
}
