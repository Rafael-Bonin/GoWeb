package handler

type SoftUpdateRequestDTO struct {
	Sobrenome string `json:"sobrenome" binding:"required"`
	Idade int `json:"idade" binding:"required"`
}