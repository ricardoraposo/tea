package main

import (
	"fmt"
	"os"
	"time"
)

type Timer struct {
	hour   int
	minute int
	second int
}

func convertToTimeFormat(targetDiff int) Timer {
	rest := targetDiff % 3600
	hour := targetDiff / 3600
	minute := rest / 60
	second := rest % 60
	return Timer{hour, minute, second}
}

func start() {
	userInput := os.Args[2]
	if len(userInput) < 2 {
		fmt.Println("Please pass a parameter brother")
		return
	}
	timer := parseInput(userInput)
	now := time.Now().Unix()
	target := time.Now().Add(time.Duration(timer.second) * time.Second).Add(time.Duration(timer.minute) * time.Minute).Add(time.Duration(timer.hour) * time.Hour).Unix()
	current := target - now
	fmt.Println(current)
}

func main() {
  createCacheDir()
	functionCall := os.Args[1]
	switch functionCall {
	case "start":
		start()
	case "help":
		fmt.Println("Nigga don't ask me.")
	default:
		fmt.Println("Invalid parameter, type 'koo help' for assistance")
	}
}
