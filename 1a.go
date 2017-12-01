package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		first int
		input string
		prev  int
		total int
	)

	_, _ = fmt.Scanln(&input)
	a := []rune(input)
	for i, r := range a {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}
		if i == 0 {
			first = d
		}
		if prev > 0 && prev == d {
			total += d
		}
		if i == len(input)-1 && d == first {
			total += d
		}
		prev = d
	}
	fmt.Println(total)
}
