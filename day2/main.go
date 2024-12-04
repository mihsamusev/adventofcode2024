package main

import (
	"bufio"
	"common"
	"fmt"
	"os"
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

	for scanner.Scan() {
		if maxScans != -1 && i == maxScans {
			break
		}
		line := scanner.Text()
		slice, err := common.ParseSlice(line)
		if err != nil {
			fmt.Printf("you're cooked: %s\n", line)
		}

		if isSafe(slice) {
			fmt.Printf("safe: %s\n", line)
			total++
		} else {
			fmt.Printf("unsafe: %s\n", line)
		}

		i++
	}

	fmt.Printf("total = %d\n", total)

	defer file.Close()
}

func isSafe(sequence []int) bool {

	hasDirection := false
	isIncreasing := sequence[1] > sequence[0]
	for i := 0; i < len(sequence) - 1; i++ {
		this := sequence[i]	
		next := sequence[i + 1]
		step := next - this
		
		if this == next {
			return false
		}

		if isIncreasing && (step < 1  || step > 3) {
			return false	
		}

		if !isIncreasing && (step > -1  || step < -3) {
			return false	
		}

		fmt.Printf("hasDirection: %t, isIncreasing: %t, step: %d\n", hasDirection, isIncreasing, step)

	}
	return true
}

