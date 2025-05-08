package dto

import (
	"github.com/google/uuid"
)

type PeopleDTO struct {
	Name        string `json:"name" validate:"required"`
	LastName    string `json:"last_name" validate:"required"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age" validate:"required,gt=0"`
	Nationalize string `json:"nationalize" validate:"required"`
	Gender      string `json:"gender" validate:"required"`
}

type PeopleFullDTO struct {
	ID uuid.UUID `json:"id"`
	PeopleDTO
}

type PeopleFilterDTO struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age"`
	Nationalize string `json:"nationalize"`
	Gender      string `json:"gender"`
}

type PeopleIdDTO struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type CreatePeopleDTO struct {
	Name       string `json:"name" validate:"required,min=2"`
	LastName   string `json:"last_name" validate:"required,min=2"`
	Patronymic string `json:"patronymic"`
}

type GetPeoplesDTO struct {
	Name        string `json:"name" query:"name"`
	LastName    string `json:"last_name" query:"last_name"`
	Patronymic  string `json:"patronymic" query:"patronymic"`
	Age         int    `json:"age" query:"age"`
	Nationalize string `json:"nationalize" query:"nationalize"`
	Gender      string `json:"gender" query:"gender"`
	Page        int    `json:"page" query:"page" validate:"gte=0"`
	Limit       int    `json:"limit" query:"limit" validate:"gte=0,lte=100"`
}
