package controller

import (
	"net/http"
	"todolist-go/modules/entities"
	"todolist-go/modules/todolist/usecase"

	"github.com/gin-gonic/gin"
)

func GetTodo(ctx *gin.Context) {
	outputRes, err := usecase.GetTodo()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"todo": outputRes})
}

func CreateTodo(ctx *gin.Context) {

	var inputReq entities.CreateTodoListReq

	if err := ctx.ShouldBindJSON(&inputReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	outputRes, err := usecase.CreateTodo(&inputReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"todo": outputRes})
}

func EditTodo(ctx *gin.Context) {
	var inputReq entities.TodoEditReq

	if err := ctx.ShouldBindJSON(&inputReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.EditTodo(&inputReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"todo": "edit success"})
}

func DelTodo(ctx *gin.Context) {
	var inputReq entities.TodoDelReq

	if err := ctx.ShouldBindJSON(&inputReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := usecase.DelTodo(&inputReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"todo": "del success"})
}
