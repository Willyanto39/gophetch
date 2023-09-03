package main

import (
	"fmt"
	"os"
	"strings"
)

func DisplaySystemInformation() {
	displayComputerModelInformation()
	displayDistroInformation()
	displayKernelInformation()
	displayCpuInformation()
	displayMemoryInformation()
	displayDesktopEnvironmentInformation()
	displayShellInformation()
}

func displayComputerModelInformation() {
	modelInformation, err := os.ReadFile("/sys/devices/virtual/dmi/id/product_version")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	display("Model", strings.Replace(string(modelInformation), "\n", "", -1))
}

func displayDistroInformation() {
	osInformation, err := os.ReadFile("/etc/os-release")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	for _, v := range strings.Split(string(osInformation), "\n") {
		keyValuePair := strings.Split(v, "=")

		if keyValuePair[0] == "PRETTY_NAME" {
			display("OS", strings.Replace(keyValuePair[1], "\"", "", -1))
			break
		}
	}
}

func displayKernelInformation() {
	kernelInformation, err := os.ReadFile("/proc/version")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	informationList := strings.Split(string(kernelInformation), " ")

	display("Kernel", informationList[0] + " " + informationList[2])
}

func displayCpuInformation() {
	cpuInformation, err := os.ReadFile("/proc/cpuinfo")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	for _, v := range strings.Split(string(cpuInformation), "\n") {
		keyValuePair := strings.Split(v, ":")

		if strings.TrimSpace(keyValuePair[0]) == "model name" {
			display("CPU", strings.TrimSpace(keyValuePair[1]))
			break
		}
	}
}

func displayMemoryInformation() {
	memoryInformation, err := os.ReadFile("/proc/meminfo")

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(0)
	}

	for _, v := range strings.Split(string(memoryInformation), "\n") {
		keyValuePair := strings.Split(v, ":")

		if keyValuePair[0] == "MemTotal" {
			display("Memory", strings.TrimSpace(keyValuePair[1]))
			break
		}
	}
}

func displayDesktopEnvironmentInformation() {
	value := getEnvironmentValue("XDG_CURRENT_DESKTOP")
	display("DE", value)
}

func displayShellInformation() {
	value := getEnvironmentValue("SHELL")
	display("Shell", value)
}

func getEnvironmentValue(environment string) string {
	value := os.Getenv(environment)

	// split value by "/" in case of getting shell name
	splitValues := strings.Split(value, "/")

	return splitValues[len(splitValues) - 1]
}

func display(label, value string) {
	fmt.Println(fmt.Sprintf("\x1b[1;37m%s\x1b[0m: %s", label, value))
}