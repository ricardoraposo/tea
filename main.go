package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Timer struct {
	hour   time.Duration
	minute time.Duration
	second time.Duration
}

func show(path string) {
	pathToVars := filepath.Join(path, "vars")
	target := readVar(pathToVars)
	now := int(time.Now().Unix())
	current := target - now
	timer := convertToTimeFormat(current)
  printToFormat(timer)
}

func start(path string) {
	userInput := os.Args[2]
	if len(userInput) < 2 {
		fmt.Println("Please pass a parameter brother")
		return
	}
	timer := parseInput(userInput)
	target := time.Now().Add(timer.second * time.Second).Add(timer.minute * time.Minute).Add(timer.hour * time.Hour).Unix()
	pathToVars := filepath.Join(path, "vars")
	err := os.WriteFile(pathToVars, []byte("duration="+strconv.FormatInt(target, 10)), 0644)
	check(err)
}

func main() {
	cachePath, _ := createCacheDir()
	functionCall := os.Args[1]
	switch functionCall {
	case "start":
		start(cachePath)
	case "show":
		show(cachePath)
	case "help":
		fmt.Println("Don't ask me")
	default:
		fmt.Println("Invalid parameter, type 'tea help' for assistance")
	}
}
