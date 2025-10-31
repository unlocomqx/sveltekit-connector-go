// Code generated. DO NOT EDIT.

package routes

import "gokit/routes/users"

// FunctionRegistry maps function names to their implementations
// TODO: Generate using the plugin
var FunctionRegistry = map[string]any{
	"src/routes/todos.remote.go:QueryTodos": func(postData map[string]any) (any, error) {
		return QueryTodos(postData)
	},
	"src/routes/todos.remote.go:FormCreateTodo": func(postData map[string]any) (any, error) {
		return FormCreateTodo(postData)
	},
	"src/routes/users/users.remote.go:QueryUserInfo": func(postData map[string]any) (any, error) {
		return users.QueryUserInfo()
	},
}
