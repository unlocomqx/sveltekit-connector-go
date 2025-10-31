package main

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: 1, Title: "Todo 1"},
	{ID: 2, Title: "Todo 2"},
	{ID: 3, Title: "Todo 3"},
}

func queryTodos() []Todo {
	return todos
}

func formCreateTodo(title string) []Todo {
	newTodo := Todo{
		ID:    len(todos) + 1,
		Title: title,
	}
	todos = append(todos, newTodo)
	return todos
}
