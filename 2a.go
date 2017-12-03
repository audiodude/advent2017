package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	r, err := regexp.Compile("(\\d+)\\s+")
	if err != nil {
		panic(err)
	}

	row_checksums := make([]int, 0)
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

		max := -1
		min := math.MaxInt32
		for _, v := range values {
			d, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}
			if d < min {
				min = d
			}
			if d > max {
				max = d
			}
		}
		row_checksums = append(row_checksums, max-min)
	}
	var sum int
	for _, c := range row_checksums {
		sum += c
	}
	fmt.Println(sum)
}
