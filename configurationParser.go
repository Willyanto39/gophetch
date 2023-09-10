package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Informations []string
}

func LoadConfiguration() Configuration {
	configuration, readErr := os.ReadFile("gophetch.json")

	if readErr != nil {
		fmt.Println("Error:", readErr)
		os.Exit(0)
	}

	var parseResult Configuration

	parseErr := json.Unmarshal(configuration, &parseResult)

	if parseErr != nil {
		fmt.Print("Error:", parseErr)
		os.Exit(0)
	}

	return parseResult
}
