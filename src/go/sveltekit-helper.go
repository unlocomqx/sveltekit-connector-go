package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"gokit/routes"
)

func executeRemoteFunction(filePath string, functionName string, postData []byte) ([]byte, error) {
	fn, exists := routes.FunctionRegistry[functionName]
	if !exists {
		return nil, fmt.Errorf("function %s not found in registry", functionName)
	}

	fnValue := reflect.ValueOf(fn)
	fnType := fnValue.Type()

	var results []reflect.Value

	if fnType.NumIn() == 0 {
		results = fnValue.Call([]reflect.Value{})
	} else if fnType.NumIn() == 1 {
		if len(postData) == 0 {
			return nil, fmt.Errorf("function %s requires parameters but none provided", functionName)
		}

		paramType := fnType.In(0)
		paramValue := reflect.New(paramType)

		if err := json.Unmarshal(postData, paramValue.Interface()); err != nil {
			return nil, fmt.Errorf("failed to unmarshal parameters: %w", err)
		}

		results = fnValue.Call([]reflect.Value{paramValue.Elem()})
	} else {
		return nil, fmt.Errorf("function %s has unsupported number of parameters", functionName)
	}

	if len(results) == 0 {
		return json.Marshal(nil)
	}

	result := results[0].Interface()
	return json.Marshal(result)
}
