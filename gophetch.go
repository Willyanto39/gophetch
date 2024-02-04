package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/Willyanto39/gophetch/config"
	"github.com/Willyanto39/gophetch/system"
)

func checkOs() {
	if runtime.GOOS != "linux" {
		fmt.Println("This program is only compatible with linux")
		os.Exit(0)
	}
}

func main() {
	checkOs()
	config := config.LoadConfiguration()
	systemInformations := system.GetSystemInformation(config)
	for _, systemInformation := range systemInformations {
		fmt.Println(systemInformation)
	}
}
