package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	dataFile := "test.txt"
	maxScans := -1
	args := os.Args
	if len(args) > 1 {
		dataFile = args[1]
	}

	if len(args) > 2 {
		maxScans, _ = strconv.Atoi(args[2])
	}

	fmt.Printf("Analyzing %d lines of %s\n", maxScans, dataFile)
	parse(dataFile, maxScans)
}

func parse(dataFile string, maxScans int) {
	file, err := os.Open(dataFile)
	if err != nil {
		fmt.Println("im dead")
	}
	scanner := bufio.NewScanner(file)

	total := 0
	i := 0

	lows := make([]int, 0)
	highs := make([]int, 0)
	for scanner.Scan() {
		if maxScans != -1 && i == maxScans {
			break
		}
		line := scanner.Text()
		slice, err:= common.ParseSlice(line)
		if err != nil {
			fmt.Printf("gay -> %s\n", line)
		}
		lows = append(lows, slice[0])
		highs = append(highs, slice[1])



		i++
	}

	sort.Ints(lows)
	sort.Ints(highs)

	//total = computeSimilarityOne(lows, highs)
	total = computeSimilarityTwo(lows, highs)
	fmt.Printf("total = %d\n", total)

	defer file.Close()
}

func computeSimilarityOne(left, right []int) int {
	total := 0
	for i := range left {
		dist := left[i] - right[i]
		if dist > 0 {
			total += dist
		} else {
			total -= dist
		}
	}
	return total
}

func computeSimilarityTwo(left, right []int) int {
	last := len(right) - 1
	pointer := 0
	total := 0
	cache := make(map[int]int, 0)

	for _, value := range left {
		_, exists := cache[value]
		if !exists {
			for (right[pointer] != value && pointer != last) {
				pointer = common.Min(pointer + 1, last)
			}

			for right[pointer] == value {
				cache[value] += value
				pointer = common.Min(pointer + 1, last)
				if pointer == last {
					break
				}
			}
			
			if pointer == last {
				pointer = 0
			}
		}

		similarity := cache[value]
		total += similarity
	}

	return total
}
