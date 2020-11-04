package apiserver

import (
	"bytes"
	"docs/internal/model"
	"docs/internal/store/teststore"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

var (
	doc1 = model.Document{
		Name:   "doc1.doc",
		Date:   "20200101",
		Number: 1,
		Sum:    strings.Repeat("a", 64),
	}
	doc2 = model.Document{
		Name:   "doc2.doc",
		Date:   "20200202",
		Number: 2,
		Sum:    strings.Repeat("b", 64),
	}
	doc3 = model.Document{
		Name:   "doc3.doc",
		Date:   "20200303",
		Number: 3,
		Sum:    strings.Repeat("c", 64),
	}
	doc4 = model.Document{
		Name:   "doc4.doc",
		Date:   "20200404",
		Number: 4,
		Sum:    strings.Repeat("d", 64),
	}
)

func marshal(data map[string][]model.Document) io.Reader {
	b, _ := json.Marshal(data)
	return bytes.NewReader(b)
}

func TestServer_handleSet(t *testing.T) {
	config := NewConfig()

	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	store := teststore.New()
	server := newServer(config, logger, store)

	tests := []struct {
		name   string
		data   io.Reader
		status int
	}{
		{
			name: "Adding doc1",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc1,
				},
			}),
			status: http.StatusOK,
		},
		{
			name: "Invalid body(0 documents)",
			data: marshal(map[string][]model.Document{
				"documents": {},
			}),
			status: http.StatusUnprocessableEntity,
		},
		{
			name: "Invalid body(empty doc name)",
			data: marshal(map[string][]model.Document{
				"documents": {
					{
						Name:   "",
						Date:   "20201010",
						Number: 1,
						Sum:    strings.Repeat("a", 64),
					},
				},
			}),
			status: http.StatusUnprocessableEntity,
		},
		{
			name:   "Invalid body(empty data)",
			data:   nil,
			status: http.StatusBadRequest,
		},
		{
			name: "Re-adding doc1",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc1,
				},
			}),
			status: http.StatusUnprocessableEntity,
		},
		{
			name: "Adding doc1 & doc2",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc1,
					doc2,
				},
			}),
			status: http.StatusUnprocessableEntity,
		},
		{
			name: "Adding doc1 & doc2 & doc3",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc1,
					doc2,
					doc3,
				},
			}),
			status: http.StatusUnprocessableEntity,
		},
		{
			name: "Adding doc2 & doc3",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc2,
					doc3,
				},
			}),
			status: http.StatusOK,
		},
		{
			name:   "Invalid JSON",
			data:   strings.NewReader("{invalid json"),
			status: http.StatusBadRequest,
		},
		{
			name: "Adding doc4",
			data: marshal(map[string][]model.Document{
				"documents": {
					doc4,
				},
			}),
			status: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodPost, "/api/document/set", tt.data)

			server.ServeHTTP(recorder, request)
			assert.Equal(t, tt.status, recorder.Code)
		})
	}
}

func TestServer_handleGet(t *testing.T) {
	config := NewConfig()

	logger := logrus.New()
	logger.SetOutput(ioutil.Discard)

	store := teststore.New()
	server := newServer(config, logger, store)

	tests := []struct {
		name   string
		status int
	}{
		{
			name:   "Get all documents",
			status: http.StatusOK,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()

			request, _ := http.NewRequest(http.MethodGet, "/api/document/get", nil)

			server.ServeHTTP(recorder, request)
			assert.Equal(t, tt.status, recorder.Code)
		})
	}
}
