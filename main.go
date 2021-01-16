package main

import (
	"caixa-eletronico/controller"
	"caixa-eletronico/tasks"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/cors"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()

	tasks.Migrate()

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"PUT", "GET", "DELETE", "POST"},
		AllowHeaders:    []string{"Content-type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length", "Content-type"},
		MaxAge:          36000,
	}))

	router.POST("/user", controller.CreateUser)
	router.PUT("/user/:id", controller.UpdateUser)
	router.GET("/user/:id", controller.ShowUser)
	router.GET("/user", controller.IndexUsers)
	router.DELETE("/user/:id", controller.DeleteUser)

	router.POST("/deposit", controller.Depositar)
	router.POST("/withdraw", controller.Sacar)

	router.Run()
}
