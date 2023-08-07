package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	h "github.com/ricardoraposo/tea/helpers"
)

func show(path string) {
	pathToVars := filepath.Join(path, "vars")
	target := h.ReadVar(pathToVars)
  if target == 0 {
    fmt.Println("")
    return
  }
	now := int(time.Now().Unix())
	current := target - now
	timer := h.ConvertToTimeFormat(current)
	h.PrintToFormat(timer)
}

func start(path string) {
	var userInput string
	if len(os.Args) < 3 {
		userInput = "30m42s"
	} else {
		userInput = os.Args[2]
	}
	timer := h.ParseInput(userInput)
	target := time.Now().Add(timer.Second * time.Second).Add(timer.Minute * time.Minute).Add(timer.Hour * time.Hour).Unix()
	pathToVars := filepath.Join(path, "vars")
	err := os.WriteFile(pathToVars, []byte("duration="+strconv.FormatInt(target, 10)), 0644)
	if err != nil {
		return
	}
}

func stop(path string) {
	pathToVars := filepath.Join(path, "vars")
	err := os.WriteFile(pathToVars, []byte(""), 0644)
  if err != nil {
    return
  }
}

func main() {
	cachePath, _ := h.CreateCacheDir()
	var functionCall string
	if len(os.Args) < 2 {
		functionCall = ""
	} else {
		functionCall = os.Args[1]
	}
	switch functionCall {
	case "start":
		start(cachePath)
	case "stop":
		stop(cachePath)
	case "show":
		show(cachePath)
	case "help":
		fmt.Println("Don't ask me")
	default:
		show(cachePath)
	}
}
