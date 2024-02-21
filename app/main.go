package main

import (
	"todolist-go/modules/server"
	"todolist-go/modules/todolist/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	server.MigrationDB()
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/todolist", controller.GetTodo)
	r.POST("/create", controller.CreateTodo)
	r.PATCH("/edit", controller.EditTodo)
	r.PATCH("/delete", controller.DelTodo)

	r.Run("localhost:8008")
}
