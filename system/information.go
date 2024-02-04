package system

import (
	"fmt"
	"os"
	"strings"

	"github.com/Willyanto39/gophetch/config"
)

var informationToOutput = map[string]string{
	"computer_model":      formatOutput("Computer Model", getComputerModel()),
	"distro":              formatOutput("Distro", getDistro()),
	"kernel":              formatOutput("Kernel", getKernel()),
	"cpu":                 formatOutput("CPU", getCpu()),
	"memory":              formatOutput("Memory", getMemory()),
	"shell":               formatOutput("Shell", getShell()),
	"desktop_environment": formatOutput("Desktop Environment", getDesktopEnvironment()),
}

func GetSystemInformation(configuration config.Configuration) []string {
	systemInformations := []string{}

	for _, information := range configuration.Informations {
		systemInformations = append(systemInformations, informationToOutput[information])
	}

	return systemInformations
}

func getComputerModel() string {
	computerModel := readFileContent("/sys/devices/virtual/dmi/id/product_version")

	return strings.Replace(computerModel, "\n", "", -1)
}

func getDistro() string {
	distro := readFileContent("/etc/os-release")

	for _, v := range strings.Split(distro, "\n") {
		keyValuePair := strings.Split(v, "=")

		if keyValuePair[0] == "PRETTY_NAME" {
			return strings.Replace(keyValuePair[1], "\"", "", -1)
		}
	}

	return "Unknown Distro"
}

func getKernel() string {
	kernel := readFileContent("/proc/version")
	kernelInformations := strings.Split(kernel, " ")

	return kernelInformations[0] + " " + kernelInformations[2]
}

func getCpu() string {
	cpu := readFileContent("/proc/cpuinfo")

	for _, v := range strings.Split(cpu, "\n") {
		keyValuePair := strings.Split(v, ":")

		if strings.TrimSpace(keyValuePair[0]) == "model name" {
			return strings.TrimSpace(keyValuePair[1])
		}
	}

	return "Unknown CPU"
}

func getMemory() string {
	memory := readFileContent("/proc/meminfo")

	for _, v := range strings.Split(memory, "\n") {
		keyValuePair := strings.Split(v, ":")

		if keyValuePair[0] == "MemTotal" {
			return strings.TrimSpace(keyValuePair[1])
		}
	}

	return "Memory Information Unknown"
}

func getShell() string {
	shell := getEnvironmentValue("SHELL")
	paths := strings.Split(shell, "/")

	return paths[len(paths)-1]
}

func getDesktopEnvironment() string {
	return getEnvironmentValue("XDG_CURRENT_DESKTOP")
}

func readFileContent(filePath string) string {
	fileContent, err := os.ReadFile(filePath)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	return string(fileContent)
}

func getEnvironmentValue(environment string) string {
	return os.Getenv(environment)
}

func formatOutput(label string, value string) string {
	return fmt.Sprintf("\x1b[1;37m%-20s\x1b[0m: %s", label, value)
}
