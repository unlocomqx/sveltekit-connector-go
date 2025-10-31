package routes

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

var todos = []Todo{
	{ID: 1, Title: "Todo 1"},
	{ID: 2, Title: "Todo 2"},
	{ID: 3, Title: "Todo 3"},
}

func QueryTodos(postData map[string]any) (any, error) {
	return todos, nil
}

func FormCreateTodo(postData map[string]any) (any, error) {
	title := postData["title"].(string)
	newTodo := Todo{
		ID:    len(todos) + 1,
		Title: title,
	}
	todos = append(todos, newTodo)
	return todos, nil
}

func FormDeleteTodo(postData map[string]any) (any, error) {
	id := postData["id"].(int)
	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return todos, nil
		}
	}
	return todos, nil
}
