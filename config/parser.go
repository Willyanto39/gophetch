package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Informations []string
}

func LoadConfiguration() Configuration {
	configFile, err := os.ReadFile("gophetch.json")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	var config Configuration
	err = json.Unmarshal(configFile, &config)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	return config
}
