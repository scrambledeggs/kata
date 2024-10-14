package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

// YAML structure:
// paths:
//   /v1/merchants/{merchant_id}/stores:
//     post:
//       [more here]
//   /v1/merchants:
//     get:
//       [more here]

const DEFAULT_PATH = ".aws-sam/build/AllEndpoints/endpoints.yml"

type Config struct {
	Paths map[string]map[string]interface{} `yaml:"paths"`
}

func (c *Config) getConf() *Config {
	file, err := os.ReadFile("docs/api_contract.yaml")
	if err != nil {
		panic(fmt.Sprintf("Error reading file: %v", err))
	}

	err = yaml.Unmarshal(file, c)
	if err != nil {
		panic(fmt.Sprintf("Error unmarshalling YAML: %v", err))
	}

	return c
}

func main() {
	var c Config
	resources := c.getConf().Paths
	endpoints := make([]string, 0)

	for path, methods := range resources {
		for method := range methods {
			endpoint := fmt.Sprintf("%s %s", strings.ToUpper(method), path)
			fmt.Println(endpoint)

			endpoints = append(endpoints, endpoint)
		}
	}

	yamlData, err := yaml.Marshal(&endpoints)
	if err != nil {
		panic(fmt.Sprintf("Error while marshaling YAML: %v", err))
	}

	path := flag.String("path", DEFAULT_PATH, "File Path")
	flag.Parse()

	err = os.WriteFile(*path, yamlData, 0644)
	if err != nil {
		panic("Unable to write data into the file")
	}

	fmt.Println("Successfully generated endpoints.yml")
}
