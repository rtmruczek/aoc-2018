package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	datapath := os.Getenv("PATH_TO_DATA_FILE")

	b, err := ioutil.ReadFile(datapath)
	if err != nil {
		panic(err)
	}
	s := string(b)
	parseFile(s)
}

func parseFile(body string) {
	lines := strings.Split(body, "\n")
	total := 0

	foundFirstFrequencyReachedTwice := false
	foundFirstTotal := false
	frequenciesReached := make(map[int]int)

	for !foundFirstFrequencyReachedTwice {
		for _, line := range lines {
			nextValue := getValueFromLine(line)
			total += nextValue
			_, isPresent := frequenciesReached[total]
			if isPresent {
				println(total)
				foundFirstFrequencyReachedTwice = true
				break
			} else {
				frequenciesReached[total] = 1
			}
		}
		if !foundFirstTotal {
			println(total)
		}
		foundFirstTotal = true

	}
}

func getValueFromLine(line string) int {
	operator := line[0]
	value, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(err)
	}
	if operator == 45 {
		return -value
	}
	return value
}
