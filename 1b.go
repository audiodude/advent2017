package main

import (
	"fmt"
	"strconv"
)

func main() {
	var (
		input string
		total int
	)

	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	l := len(input) / 2
	double := []rune(input + input)
	a := []rune(input)
	for i, r := range a {
		d, err := strconv.Atoi(string(r))
		if err != nil {
			panic(err)
		}

		e, err := strconv.Atoi(string(double[i+l]))
		if err != nil {
			panic(err)
		}
		if d == e {
			total += d
		}
	}
	fmt.Println(total)
}
