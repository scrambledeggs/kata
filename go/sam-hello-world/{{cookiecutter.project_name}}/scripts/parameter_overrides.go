package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	filePath := ".secrets.json"
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	jsonFile, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}

	var params map[string]string
	if err := json.Unmarshal(byteValue, &params); err != nil {
		log.Fatalf("Failed to parse JSON: %s", err)
	}

	var result []string
	for key, value := range params {
		result = append(result, fmt.Sprintf("ParameterKey=%s,ParameterValue=%s", key, value))
	}

	fmt.Println(strings.Join(result, " "))
}
