package main

import (
	"regexp"
	"strconv"
)

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
	return Timer{hour, minute, second}
}

func convertToTimeFormat(targetDiff int) Timer {
	rest := targetDiff % 3600
	hour := targetDiff / 3600
	minute := rest / 60
	second := rest % 60
	return Timer{hour, minute, second}
}
