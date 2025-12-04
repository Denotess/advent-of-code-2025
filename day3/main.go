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
	res := 0
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var text string = string(scanner.Text())
		//fmt.Println(string(text))
		max := -1
		for i := 0; i < len(text); i++ {
			for j := i + 1; j < len(text); j++ {
				first, err1 := strconv.Atoi(string(text[i]))
				second, err2 := strconv.Atoi(string(text[j]))
				if err1 != nil || err2 != nil {
					continue
				}
				num := first*10 + second // first*10 so no need to convert as first number is the tens place and second is ones
				if num > max {
					max = num
				}
			}
		}
		res += max
	}
	fmt.Println(res)
}
