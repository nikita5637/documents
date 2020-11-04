package model

import (
	"errors"
	"regexp"

	validation "github.com/go-ozzo/ozzo-validation"
)

//Document основная структура документа
type Document struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Number int    `json:"number"`
	Sum    string `json:"sum"`
}

//Validate проверяет валидность данных в структуре и возвращает isValid и ошибку
func (d *Document) Validate() error {
	if err := validation.ValidateStruct(d,
		validation.Field(&d.Name, validation.Required),
		validation.Field(&d.Date, validation.Required),
		validation.Field(&d.Number, validation.Required),
		validation.Field(&d.Sum, validation.Required),
	); err != nil {
		return err
	}

	checkSumRegExp, _ := regexp.Compile("^[a-fA-F0-9]{64}$")
	if checkSumRegExp.Match([]byte(d.Sum)) != true {
		return errors.New("Checksum is not valid sha256 sum")
	}

	return nil
}

//New создаёт новую структуру документа и возвращает указатель на неё
func New(name, date string, number int, sum string) (*Document, error) {
	d := &Document{
		Name:   name,
		Date:   date,
		Number: number,
		Sum:    sum,
	}

	if err := d.Validate(); err != nil {
		return nil, err
	}
	return d, nil
}
