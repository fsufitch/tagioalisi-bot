package main

import (
	"fmt"
	"os"
	"time"

	"github.com/fsufitch/discord-boar-bot/common"
)

func main() {
	runtime, err := InitializeCLIRuntime()

	if err == nil {
		runtime.Logger.Log(common.LogDebug, "hello debug")
		runtime.Logger.Log(common.LogInfo, "hello info")
		runtime.Logger.Log(common.LogWarning, "hello warn")
		runtime.Logger.Log(common.LogError, "hello error")
		<-time.After(1 * time.Second)
	} else {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
