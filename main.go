package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Timer struct {
	hour   int
	minute int
	second int
}

func start(path string) {
	userInput := os.Args[2]
	if len(userInput) < 2 {
		fmt.Println("Please pass a parameter brother")
		return
	}
	timer := parseInput(userInput)
	now := time.Now().Unix()
	target := time.Now().Add(time.Duration(timer.second) * time.Second).Add(time.Duration(timer.minute) * time.Minute).Add(time.Duration(timer.hour) * time.Hour).Unix()
	current := target - now
  err := os.WriteFile(filepath.Join(path, "vars"), []byte(strconv.FormatInt(current, 10)), 0644)
  if err != nil {
    return
  }
}

func main() {
	cachePath, _ := createCacheDir()
  fmt.Println(cachePath)
	functionCall := os.Args[1]
	switch functionCall {
	case "start":
		start(cachePath)
	case "help":
		fmt.Println("Don't ask me")
	default:
		fmt.Println("Invalid parameter, type 'tea help' for assistance")
	}
}
