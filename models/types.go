package models

type ServerResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Todo struct {
	ID          int    `json:"id"`
	Body        string `json:"body"`
	IsCompleted bool   `json:"isCompleted"`
}
