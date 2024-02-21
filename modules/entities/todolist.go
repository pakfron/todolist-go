package entities

type CreateTodoListReq struct {
	Todo string `json:"todo" binding:"required"`
}

type CreateTodoListRes struct {
	ID    int
	Todo  string
	Check bool
	IsDel bool
}

type TodolistRes struct {
	ID    int
	Todo  string
	Check bool
	IsDel bool
}

type TodoEditReq struct {
	ID   int    `json:"id" binding:"required"`
	Todo string `json:"todo" binding:"required"`
}

type TodoEditDB struct {
	Todo string
}
type IDEditDB struct {
	ID int
}

type TodoDelReq struct {
	ID    int  `json:"id" binding:"required"`
	IsDel bool `json:"is_del" binding:"required"`
}
