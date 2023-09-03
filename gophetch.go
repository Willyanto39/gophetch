package main

import (
	"fmt"
	"os"
	"runtime"
)

func checkOs() {
	if runtime.GOOS != "linux" {
		fmt.Println("This program is only compatible with linux")
		os.Exit(0)
	}
}

func main() {
	checkOs()
	DisplaySystemInformation()
}