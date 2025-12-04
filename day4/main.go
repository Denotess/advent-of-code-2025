package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(input)
	var numArray [][]string
	for scanner.Scan() {
		text := scanner.Text()
		curStrings := strings.Split(text, "")
		numArray = append(numArray, curStrings)
	}
	res := 0
	can, numArray, output := canBeRemoved(numArray)
	for can {
		res += output
		can, numArray, output = canBeRemoved(numArray)
	}
	fmt.Println(res)
}

func canBeRemoved(numArray [][]string) (bool, [][]string, int) {
	dirs := [][2]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	res := 0
	for i := 0; i < len(numArray); i++ {
		for j := 0; j < len(numArray[i]); j++ {
			counter := 0
			for _, d := range dirs {
				ni, nj := i+d[0], j+d[1]
				if ni >= 0 && ni < len(numArray) && nj >= 0 && nj < len(numArray[ni]) {
					if numArray[ni][nj] == "@" {
						counter += 1
					}
				}
			}
			if numArray[i][j] == "@" && counter < 4 {
				numArray[i][j] = "x"
				res += 1
			}
		}
	}
	for i := 0; i < len(numArray); i++ {
		for j := 0; j < len(numArray[i]); j++ {
			if numArray[i][j] == "x" {
				numArray[i][j] = "."
			}
		}
	}
	if res >= 1 {
		return true, numArray, res
	}
	return false, nil, 0
}
