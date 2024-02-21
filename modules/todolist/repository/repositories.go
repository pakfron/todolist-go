package repositories

import (
	"todolist-go/modules/entities"
	"todolist-go/modules/server"
	"todolist-go/pkg/database"
)

func GetTodolist() (*[]entities.TodolistRes, error) {
	db := server.GetDB()
	defer server.CloseDB(db)
	outputRes := []entities.TodolistRes{}
	result := db.Model(database.Todolist{}).Find(&outputRes)
	if result.Error != nil {
		return nil, result.Error
	}
	return &outputRes, nil

}

func NewTaskTodoDB(inputReq *entities.CreateTodoListReq) (*entities.CreateTodoListRes, error) {

	db := server.GetDB()
	defer server.CloseDB(db)
	todoReq := database.Todolist{
		Todo: inputReq.Todo,
	}

	result := db.Model(database.Todolist{}).Create(&todoReq)
	if result.Error != nil {
		return nil, result.Error
	}
	todoRes := entities.CreateTodoListRes{
		ID:    todoReq.ID,
		Todo:  todoReq.Todo,
		Check: todoReq.Check,
		IsDel: todoReq.IsDel,
	}
	return &todoRes, nil

}

func EditTodo(todoReq *entities.TodoEditReq) error {
	db := server.GetDB()
	defer server.CloseDB(db)

	todoEditDB := database.Todolist{
		Todo: todoReq.Todo,
	}
	result := db.Model(database.Todolist{}).Where("id = ?", todoReq.ID).Update("Todo", todoEditDB.Todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DelTodo(todoReq *entities.TodoDelReq) error {
	db := server.GetDB()
	defer server.CloseDB(db)
	result := db.Model(database.Todolist{}).Where("id = ?", todoReq.ID).Update("is_del", todoReq.IsDel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
