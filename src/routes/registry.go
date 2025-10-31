// Code generated. DO NOT EDIT.

package routes

import "encoding/json"

// FunctionRegistry maps function names to their implementations
var FunctionRegistry = map[string]any{
	"src/routes/todos.remote.go:queryTodos": func(postData []byte) any {
		return queryTodos()
	},
	"src/routes/todos.remote.go:formCreateTodo": func(postData []byte) {
		var params []string
		if err := json.Unmarshal(postData, &params); err != nil {
			return nil, err
		}
		return formCreateTodo(postData)
	},
}
