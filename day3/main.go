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
		text := scanner.Text()
		//fmt.Println(string(text))
		// for i := 0; i < len(text); i++ {
		// 	for j := i + 1; j < len(text); j++ {
		// 		first, err1 := strconv.Atoi(string(text[i]))
		// 		second, err2 := strconv.Atoi(string(text[j]))
		// 		if err1 != nil || err2 != nil {
		// 			continue
		// 		}
		// 		num := first*10 + second // first*10 so no need to convert as first number is the tens place and second is ones
		// 		if num > max {
		// 			max = num
		// 		}
		// 	}
		// }
		largest := maxJoltage12(text)
		num, err := strconv.Atoi(largest)
		if err != nil {
			panic(err)
		}
		res += num
	}
	fmt.Println(res)
}

func maxJoltage12(s string) string {
	n := len(s)
	k := 12
	result := ""
	start := 0
	for i := 0; i < k; i++ {
		maxDigit := byte('0')
		maxIdx := start
		end := n - (k - i)
		for j := start; j <= end; j++ {
			if s[j] > maxDigit {
				maxDigit = s[j]
				maxIdx = j
			}
		}
		result += string(maxDigit)
		start = maxIdx + 1
	}
	return result
}
