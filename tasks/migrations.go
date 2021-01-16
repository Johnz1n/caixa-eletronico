package tasks

import (
	"caixa-eletronico/database"
	"caixa-eletronico/model"
)

func Migrate() {
	db, _ := database.GetConnection()
	db.AutoMigrate(&model.Usuario{}, &model.Conta{})
}
