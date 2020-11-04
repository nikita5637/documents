package model

import (
	"crypto/sha256"
	"fmt"
)

//TestDocument возвращает структуру для тестов
func TestDocument() *Document {
	h := sha256.New()
	h.Write([]byte("Test sha256 hashsum"))
	return &Document{
		Name:   "Test docname.doc",
		Date:   "20201102",
		Number: 1,
		Sum:    fmt.Sprintf("%x", h.Sum(nil)),
	}
}
