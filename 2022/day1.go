package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func min(l [3]int) int {
	min := l[0]
	var min_index int
	for i := 1; i < len(l); i++ {
		if min > l[i] {
			min = l[i]
			min_index = i
		}
	}
	return min, min_index
}

func main() {
	var top_calories [3]int

	file, err := os.Open("day1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	for scanner.Scan() {
		local_max := 0
		line := scanner.Text()

		value, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}

		if line == "\n" {
			if value < min(top_calories) {

				local_max = 0
				continue
			}
		}

		local_max += value
		fmt.Println(local_max)
	}
}
