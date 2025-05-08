package dto

import "github.com/google/uuid"

type PeopleDTO struct {
	Name        string `json:"name" binding:"required"`
	LastName    string `json:"last_name" binding:"required"`
	Patronymic  string `json:"patronymic"`
	Age         int    `json:"age" binding:"required"`
	Nationalize string `json:"nationalize" binding:"required"`
	Gender      string `json:"gender" binding:"required"`
}

type PeopleIdDTO struct {
	ID uuid.UUID `json:"id" binding:"required"`
}

type CreatePeopleDTO struct {
	Name       string `json:"name" binding:"required"`
	LastName   string `json:"last_name" binding:"required"`
	Patronymic string `json:"patronymic"`
}

type GetPeoplesDTO struct {
	PeopleDTO
	Page  int `json:"page"`
	Limit int `json:"limit"`
}
