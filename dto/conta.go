package dto

import (
	"caixa-eletronico/model"
)

type ContaRequest struct {
	IDUsuario int `json:"idUsuario" validate:"required"`
	IDConta   int `json:"idconta" validate:"required"`
	Valor     int `json:"valor" validate:"required"`
}
type MessageReply struct {
	Message string `json:"message"`
}

type DepositReply struct {
	Message string      `json:"message"`
	Conta   model.Conta `json:"conta"`
}

type SaqueReply struct {
	Message string      `json:"message"`
	Notas   []Notas     `json:"notas"`
	Conta   model.Conta `json:"conta"`
}

type Notas struct {
	Valor      int `json:"valor"`
	Quantidade int `json:"quantidade"`
}
