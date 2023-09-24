package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./files/day1/input.txt")

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	scanner := bufio.NewScanner(file)

	res := []int{}
	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) <= 0 {
			res = append(res, sum)
			sum = 0
		} else {
			calories, err := strconv.Atoi(line)
			if err != nil {
				panic(err)
			}
			sum += calories
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	max := 0
	for _, value := range res {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
