package internal

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Timer struct {
	Hour   time.Duration
	Minute time.Duration
	Second time.Duration
}

func ParseInput(input string) Timer {
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

func ConvertToTimeFormat(targetDiff int) Timer {
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

func PrintToFormat(timer Timer) {
	if timer.Second < 0 || timer.Minute < 0 || timer.Hour < 0 {
		if timer.Hour == 0 && timer.Minute == 0 {
			fmt.Printf("🍵-%d\n", -timer.Second)
		} else if timer.Hour == 0 {
			fmt.Printf("🍵-%d:%02d\n", -timer.Minute, -timer.Second)
		} else {
			fmt.Printf("🍵-%d:%02d:%02d\n", -timer.Hour, -timer.Minute, -timer.Second)
		}
	} else {
		if timer.Hour == 0 && timer.Minute == 0 {
			fmt.Printf("🍵%d\n", timer.Second)
		} else if timer.Hour == 0 {
			fmt.Printf("🍵%d:%02d\n", timer.Minute, timer.Second)
		} else {
			fmt.Printf("🍵%d:%02d:%02d\n", timer.Hour, timer.Minute, timer.Second)
		}

	}
}

func ReadVar(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
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
