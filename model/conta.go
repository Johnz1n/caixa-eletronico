package model

import (
	"caixa-eletronico/database"
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
)

type TipoConta int

const (
	ContaCorrente TipoConta = 1
	ContaPopanca  TipoConta = 2
)

type Conta struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UsuarioID   uint      `json:"usuarioID"`
	UsuarioType string    `json:"-"`
	Saldo       int       `json:"saldo"`
	TipoConta   TipoConta `json:"tipoConta" validate:"required,oneof=1,2"`
}

func (c *Conta) Hash() string {
	str := "id=" + strconv.Itoa(int(c.ID))
	data := []byte(str)
	return fmt.Sprintf("%x", md5.Sum(data))
}

func (c *Conta) Get(id, idUsuario int) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Where("id = ? and usuario_id = ?", id, idUsuario).Take(&c).Error
}

func (c *Conta) Sacar(valor int) ([]int, error) {
	if valor <= 0 || valor > c.Saldo {
		return nil, errors.New("Saldo insuficiente")
	}
	notas, err := getNotas(valor)
	if err != nil {
		return nil, nil
	}

	db, err := database.GetConnection()
	if err != nil {
		return nil, err
	}

	c.Saldo -= valor
	if err := db.Save(&c).Error; err != nil {
		return nil, err
	}
	return notas, nil

}

func (c *Conta) Depositar(valor int) error {
	if valor < 0 {
		return errors.New("Valor do deposito menor que zero")
	}

	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	c.Saldo += valor
	if err := db.Save(&c).Error; err != nil {
		return err
	}
	return nil
}

func getNotas(valor int) ([]int, error) {
	valorNotas := []int{100, 50, 20}
	valorRestante := valor
	j := 0
	notas := []int{}
	next := false
	for i := 0; i < len(valorNotas); i++ {
		if valorNotas[i] > valorRestante {
			continue
		}

		valorRestante -= valorNotas[i]
		notas = append(notas, valorNotas[i])
		if valorRestante == 0 {
			break
		}
		for {
			if valorNotas[j] > valorRestante {
				j++
				if j > 2 {
					next = true
					break
				}
				continue
			}
			valorRestante -= valorNotas[j]

			if valorRestante == 0 {
				notas = append(notas, valorNotas[j])
				break
			}
			if j < 2 && valorRestante < valorNotas[j+1] {
				valorRestante += valorNotas[j]
				j++
			} else {
				notas = append(notas, valorNotas[j])
			}
		}

		if !next {
			break
		}

		valorRestante = valor
		notas = []int{}
		next = false
		j = 0
	}

	if len(notas) == 0 {
		return notas, errors.New("Valor inválido para saque")
	}
	return notas, nil
}
