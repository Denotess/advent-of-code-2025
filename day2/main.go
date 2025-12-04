package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		res := 0
		stringArray := strings.Split(line, ",")
		for i := 0; i < len(stringArray); i++ {
			currStringArray := strings.Split(stringArray[i], "-")
			start, err := strconv.Atoi(currStringArray[0])
			end, err := strconv.Atoi(currStringArray[1])
			if err != nil {
				continue
			}
			for i := start; i <= end; i++ {
				strI := strconv.Itoa(i)
				if isInvalidID(strI) {
					res += i
				}
			}
		}
		fmt.Println(res)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func isInvalidID(s string) bool {
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		pattern := s[:l]
		valid := true
		for i := l; i < n; i += l {
			if s[i:i+l] != pattern {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}
