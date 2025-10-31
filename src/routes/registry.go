// Code generated. DO NOT EDIT.

package routes

// FunctionRegistry maps function names to their implementations
var FunctionRegistry = map[string]any{
	"src/routes/todos.remote.go:queryTodos": func(postData []byte) any {
		return queryTodos(postData)
	},
	"src/routes/todos.remote.go:formCreateTodo": func(postData []byte) {
		return formCreateTodo(postData)
	},
}
