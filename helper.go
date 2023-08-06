package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

func check(err error) error {
	if err != nil {
		return err
	}
	return nil
}

func parseInput(input string) Timer {
	hrRe := regexp.MustCompile(`(\d+)h`)
	mnRe := regexp.MustCompile(`(\d+)m`)
	seRe := regexp.MustCompile(`(\d+)s`)
	hourMatches := hrRe.FindAllStringSubmatch(input, -1)
	minMatches := mnRe.FindAllStringSubmatch(input, -1)
	secMatches := seRe.FindAllStringSubmatch(input, -1)
	hour := 0
	minute := 0
	second := 0
	for _, match := range hourMatches {
		hourInt, _ := strconv.Atoi(match[1])
		hour += hourInt
	}
	for _, match := range minMatches {
		minInt, _ := strconv.Atoi(match[1])
		minute += minInt
	}
	for _, match := range secMatches {
		secInt, _ := strconv.Atoi(match[1])
		second += secInt
	}
	return Timer{
		time.Duration(hour),
		time.Duration(minute),
		time.Duration(second),
	}
}

func convertToTimeFormat(targetDiff int) Timer {
	rest := targetDiff % 3600
	hour := targetDiff / 3600
	minute := rest / 60
	second := rest % 60
	return Timer{
		time.Duration(hour),
		time.Duration(minute),
		time.Duration(second),
	}
}

func printToFormat(timer Timer) {
	if timer.second < 0 || timer.minute < 0 || timer.hour < 0 {
		if timer.hour == 0 && timer.minute == 0 {
			fmt.Printf("🍵-%d\n", -timer.second)
		} else if timer.hour == 0 {
			fmt.Printf("🍵-%d:%02d\n", -timer.minute, -timer.second)
		} else {
			fmt.Printf("🍵-%d:%02d:%02d\n", -timer.hour, -timer.minute, -timer.second)
		}
	} else {
		if timer.hour == 0 && timer.minute == 0 {
			fmt.Printf("🍵-%d\n", -timer.second)
		} else if timer.hour == 0 {
			fmt.Printf("🍵-%d:%02d\n", -timer.minute, -timer.second)
		} else {
			fmt.Printf("🍵%d:%02d:%02d\n", timer.hour, timer.minute, timer.second)
		}

	}
}

func readVar(path string) int {
	file, err := os.Open(path)
	check(err)
	scanner := bufio.NewScanner(file)
	var value int
	for scanner.Scan() {
		line := scanner.Text()
		re := regexp.MustCompile(`duration=(\d+)`)
		match := re.FindStringSubmatch(line)
		valueStr, _ := strconv.Atoi(match[1])
		value = valueStr
	}
	return value
}
