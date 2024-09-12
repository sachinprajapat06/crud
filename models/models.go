package models

type Card struct {
	Number string `json:"number" validate:"required,len=16,numeric"`
}
