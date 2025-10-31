package main

import (
	"encoding/json"
	"fmt"

	"gokit/routes"
)

func executeRemoteFunction(path string, functionName string, postData []byte) ([]byte, error) {
	registryKey := path + ":" + functionName
	fn, exists := routes.FunctionRegistry[registryKey]
	if !exists {
		return nil, fmt.Errorf("function %s not found in registry", registryKey)
	}

	result, err := fn.(func([]byte) (any, error))(postData)
	if err != nil {
		return nil, err
	}

	return json.Marshal(result)
}
