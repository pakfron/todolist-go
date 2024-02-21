package usecase

import (
	"todolist-go/modules/entities"
	repositories "todolist-go/modules/todolist/repository"
)

func CreateTodo(inputReq *entities.CreateTodoListReq) (*entities.CreateTodoListRes, error) {
	outputRes, err := repositories.NewTaskTodoDB(inputReq)
	if err != nil {
		return nil, err
	}
	return outputRes, nil
}

func GetTodo() (*[]entities.TodolistRes, error) {

	outputRes, err := repositories.GetTodolist()
	if err != nil {
		return nil, err
	}

	return outputRes, nil
}

func EditTodo(todoReq *entities.TodoEditReq) error {

	err := repositories.EditTodo(todoReq)
	if err != nil {
		return err
	}
	return nil
}

func DelTodo(todoReq *entities.TodoDelReq) error {

	err := repositories.DelTodo(todoReq)

	if err != nil {
		return err
	}
	return nil
}
