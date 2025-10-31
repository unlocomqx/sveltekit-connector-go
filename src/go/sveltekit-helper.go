package main

import (
	"gokit/routes"
)

func executeRemoteFunction(path string, functionName string, postData map[string]any) any {
	registryKey := path + ":" + functionName
	fn, exists := routes.FunctionRegistry[registryKey]
	if !exists {
		return nil
	}
	return fn.(func(map[string]any) any)(postData)
}
