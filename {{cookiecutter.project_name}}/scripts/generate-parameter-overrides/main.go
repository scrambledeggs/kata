package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const DEFAULT_PATH = ".secrets.json"

func main() {
	path := flag.String("path", DEFAULT_PATH, "Path of File")
	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
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
