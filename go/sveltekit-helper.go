package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func executeRemoteFunction(filePath string, functionName string, postData []byte) ([]byte, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	contentStr := string(fileContent)
	contentStr = strings.Replace(contentStr, "package main", "", 1)

	hasPostData := len(postData) > 0

	var wrapperCode string
	if hasPostData {
		wrapperCode = fmt.Sprintf(`package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

%s

func main() {
	postData, _ := io.ReadAll(os.Stdin)
	result := %s(postData)
	jsonData, err := json.Marshal(result)
	if err != nil {
		return
	}
	fmt.Print(string(jsonData))
}
`, contentStr, functionName)
	} else {
		wrapperCode = fmt.Sprintf(`package main

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
	}

	wrapperPath := filepath.Join("tmp", "wrapper.go")
	err = os.WriteFile(wrapperPath, []byte(wrapperCode), 0644)
	if err != nil {
		return nil, err
	}

	cmd := exec.Command("go", "run", wrapperPath)

	if len(postData) > 0 {
		stdin, err := cmd.StdinPipe()
		if err != nil {
			return nil, err
		}
		go func() {
			defer stdin.Close()
			stdin.Write(postData)
		}()
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("execution error: %s", string(output))
	}

	return []byte(strings.TrimSpace(string(output))), nil
}
