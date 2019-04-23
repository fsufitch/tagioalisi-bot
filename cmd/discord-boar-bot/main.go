package main

import (
	"fmt"
	"os"
)

func main() {
	runtime, err := InitializeCLIRuntime()

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	err = runtime.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
