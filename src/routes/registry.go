// Code generated. DO NOT EDIT.

package routes

// FunctionRegistry maps function names to their implementations
var FunctionRegistry = map[string]any{
	"src/routes/todos.remote.go:queryTodos": func(postData map[string]any) (any, error) {
		return queryTodos(postData)
	},
	"src/routes/todos.remote.go:formCreateTodo": func(postData map[string]any) (any, error) {
		return formCreateTodo(postData)
	},
	"src/routes/users/user.remote.go:queryUserInfo": func(postData map[string]any) (any, error) {
		return fetchUserInfo(postData)
	},
}
