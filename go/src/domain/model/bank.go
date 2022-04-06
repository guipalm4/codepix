package model

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type Bank struct {
	ID        string    `json: "id"`
	Code      string    `json: "code"`
	Name      string    `json: "name"`
	CreatedAt time.Time `json: "createdAt"`
	UpdatedAt time.Time `json: "updatedAt"`
}

func NewBank(code string, name string) (*Bank, error) {
	bank := Bank{
		Code: code,
		Name: name,
	}

	bank.ID = uuid.NewV4()

	return &bank, nil
}
