package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strconv"
)

func main() {
	r, err := regexp.Compile("(\\d+)\\s+")
	if err != nil {
		panic(err)
	}

	var sum int
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		} else if err == io.EOF {
			break
		}
		match := r.FindAllStringSubmatch(input, -1)
		values := make([]string, 0)
		for _, m := range match {
			values = append(values, m[1])
		}

		values_int := make([]int, 0)
		for _, v := range values {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			values_int = append(values_int, d)
		}
		sort.Ints(values_int)

		for i := len(values_int) - 1; i >= 0; i-- {
			var found int
			big := values_int[i]
			for _, v := range values_int {
				if float64(big)/float64(v) < 2 {
					break
				}
				if big%v == 0 {
					found = big / v
					break
				}
			}
			if found > 0 {
				sum += found
				break
			}
		}
	}
	fmt.Println(sum)
}
