package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	comb := 50
	counter := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		num := line[1:]
		comb, err = calc(direction, num, comb)
		if comb == 0 {
			counter++
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(counter)
}

func calc(direction string, num string, comb int) (int, error) {
	n, err := strconv.Atoi(num)
	if err != nil {
		return comb, err
	}
	if direction == "L" {
		comb -= n
		comb = (comb + 100) % 100
	}
	if direction == "R" {
		comb += n
		comb = (comb + 100) % 100
	}
	return comb, nil
}
