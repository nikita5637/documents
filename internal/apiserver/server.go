package apiserver

import (
	"docs/internal/model"
	"docs/internal/store"
	"encoding/json"
	"errors"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

//Server структура, хранящая переменные и методы REST API сервера
type Server struct {
	config *Config
	logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func newServer(config *Config, logger *logrus.Logger, store store.Store) *Server {
	s := &Server{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/api/document/set", s.handleSet()).Methods(http.MethodPost)
	s.router.HandleFunc("/api/document/get", s.handleGet()).Methods(http.MethodGet)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleSet() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Body == nil {
			s.error(w, r, http.StatusBadRequest, errors.New("Empty request body"))
			return
		}
		defer r.Body.Close()

		var req = requestBody{
			store: s.store,
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}

		if err := req.validate(s.config.MinDocsPerRequest, s.config.MaxDocsPerRequest); err != nil {
			s.logger.Warnf("Error while validating request: %s", err)
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}

		var errorsChan = make(chan error)
		var wgDoneChan = make(chan bool)
		var wg sync.WaitGroup

		for _, doc := range req.Documents {
			wg.Add(1)
			lDoc := doc
			go func(doc *model.Document) {
				defer wg.Done()
				if err := s.store.Docs().InsertDoc(doc); err != nil {
					errorsChan <- err
					return
				}
				s.logger.Infof("Successfully added new document %+v", *doc)
			}(&lDoc)
		}

		go func() {
			wg.Wait()
			close(wgDoneChan)
		}()

		select {
		case <-wgDoneChan:
			s.respond(w, r, http.StatusOK, nil)
			return
		case err := <-errorsChan:
			s.logger.Warnf("Error while adding new document: %s", err)
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}
	}
}

func (s *Server) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		docs, err := s.store.Docs().GetDocs()
		if err != nil {
			s.error(w, r, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, r, http.StatusOK, map[string][]model.Document{
			"documents": *docs,
		})
	}
}

func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, struct {
		ErrorMessage string
	}{
		ErrorMessage: err.Error(),
	})
}

func (s *Server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
