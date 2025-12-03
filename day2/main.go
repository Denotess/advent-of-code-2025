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
	passes := 0

	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[:1]
		num := line[1:]
		comb, passes, err = calc(direction, num, comb)
		if err != nil {
			panic(err)
		}
		counter += passes

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	fmt.Println(counter)
}

func calc(direction string, num string, comb int) (int, int, error) {
	passes := 0
	n, err := strconv.Atoi(num)
	if err != nil {
		return comb, 0, err
	}
	if direction == "L" {
		for i := 1; i <= n; i++ {
			if (comb-i+100)%100 == 0 {
				passes += 1
			}

		}
		comb -= n
		comb = (comb + 100) % 100
	}
	if direction == "R" {
		for i := 1; i <= n; i++ {
			if (comb+i)%100 == 0 {
				passes += 1
			}
		}
		comb += n
		comb = (comb + 100) % 100
	}
	return comb, passes, nil
}
