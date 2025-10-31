package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func executeRemoteFunction(filePath string, functionName string) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(fileContent)
	contentStr = strings.Replace(contentStr, "package main", "", 1)

	wrapperCode := fmt.Sprintf(`package main

import (
	"encoding/json"
	"fmt"
)

%s

func main() {
	result := %s()
	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Print(string(jsonData))
}
`, contentStr, functionName)

	wrapperPath := filepath.Join("tmp", "wrapper.go")
	err = os.WriteFile(wrapperPath, []byte(wrapperCode), 0644)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "run", wrapperPath)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("execution error: %s", string(output))
	}

	return []byte(strings.TrimSpace(string(output))), nil
}
