package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	i "github.com/ricardoraposo/tea/internal"
)

func show(path string) {
	pathToVars := filepath.Join(path, "vars")
	target := i.ReadVar(pathToVars)
	now := int(time.Now().Unix())
	current := target - now
	timer := i.ConvertToTimeFormat(current)
	i.PrintToFormat(timer)
}

func start(path string) {
	if len(os.Args) < 2 {
		fmt.Println("Please pass a parameter brother")
		return
	}
	userInput := os.Args[2]
	timer := i.ParseInput(userInput)
	target := time.Now().Add(timer.Second * time.Second).Add(timer.Minute * time.Minute).Add(timer.Hour * time.Hour).Unix()
	pathToVars := filepath.Join(path, "vars")
	err := os.WriteFile(pathToVars, []byte("duration="+strconv.FormatInt(target, 10)), 0644)
	if err != nil {
		return
	}
}

func main() {
	cachePath, _ := i.CreateCacheDir()
  var functionCall string
	if len(os.Args) < 2 {
    functionCall = ""
	} else {
    functionCall = os.Args[1]
  }
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
