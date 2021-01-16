package model

import (
	"caixa-eletronico/database"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Usuario struct {
	ID             uint           `json:"id" gorm:"primaryKey"`
	Nome           string         `json:"nome" validate:"required"`
	DataNascimento time.Time      `json:"dataNascimento" validate:"required"`
	Cpf            string         `json:"cpf" validate:"required"`
	Contas         []Conta        `json:"contas" gorm:"polymorphic:Usuario"`
	CreatedAt      time.Time      `json:"createdAt"`
	UpdatedAt      time.Time      `json:"updatedAt"`
	DeletedAt      gorm.DeletedAt `json:"-"`
}

type UsuarioList []Usuario

func (u *Usuario) Create() error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	for i := range u.Contas {
		u.Contas[i].Saldo = 0
	}

	return db.Create(&u).Error
}

func (u *Usuario) Update(new Usuario) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	u.Nome = new.Nome
	u.DataNascimento = new.DataNascimento
	u.Cpf = new.Cpf
	u.Contas = new.Contas
	if err := db.Save(&u).Error; err != nil {
		return err
	}

	contasID := []uint{}
	for i := range u.Contas {
		contasID = append(contasID, u.Contas[i].ID)
	}

	if err := db.Where("usuario_id = ? and id not in ?", u.ID, contasID).Delete(&Conta{}).Error; err != nil {
		return err
	}

	return nil
}

func (u *Usuario) FindByID(id int) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Where("id = ?", id).Preload(clause.Associations).Take(&u).Error
}

func (u *Usuario) FindAll() (UsuarioList, error) {
	usuarios := UsuarioList{}
	db, err := database.GetConnection()
	if err != nil {
		return usuarios, err
	}

	if err := db.Preload(clause.Associations).Find(&usuarios).Error; err != nil {
		return usuarios, nil
	}
	return usuarios, nil
}

func (u *Usuario) Delete(id int) error {
	db, err := database.GetConnection()
	if err != nil {
		return err
	}

	return db.Where("id = ?", id).Delete(&u).Error
}
