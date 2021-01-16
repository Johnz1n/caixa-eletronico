package controller

import (
	"caixa-eletronico/dto"
	"caixa-eletronico/errors"
	"caixa-eletronico/model"
	"caixa-eletronico/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Depositar(c *gin.Context) {
	var request dto.ContaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid payload"})
		return
	}

	if err := util.Validate(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: err.Error()})
		return
	}

	var conta model.Conta
	if err := conta.Get(request.IDConta, request.IDUsuario); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to get account"})
		return
	}

	if err := conta.Depositar(request.Valor); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.DepositReply{Message: "Successfully deposited", Conta: conta})
	return
}

func Sacar(c *gin.Context) {
	var request dto.ContaRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: "Invalid payload"})
		return
	}

	if err := util.Validate(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errors.GenericError{Error: err.Error()})
		return
	}

	var conta model.Conta
	if err := conta.Get(request.IDConta, request.IDUsuario); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: "Falied to get account"})
		return
	}

	notas, err := conta.Sacar(request.Valor)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, errors.GenericError{Error: err.Error()})
		return
	}

	notasMap := map[int]int{
		20:  0,
		50:  0,
		100: 0,
	}
	notaComparer := notas[0]
	for _, nota := range notas {
		if notaComparer == nota {
			notasMap[nota]++
		} else {
			notaComparer = nota
			notasMap[nota]++
		}
	}

	notasReply := []dto.Notas{}

	for valor, quantidade := range notasMap {
		notasReply = append(notasReply, dto.Notas{Valor: valor, Quantidade: quantidade})
	}

	c.JSON(http.StatusOK, dto.SaqueReply{Message: "Successfully withdrawn", Notas: notasReply, Conta: conta})
	return
}
