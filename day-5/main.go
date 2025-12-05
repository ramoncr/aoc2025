package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	start, end int
}

func main() {
	start := time.Now()

	fmt.Println("Reading input...")
	freshRanges, productIds := readFreshRangesAndProducts()

	fmt.Println("Converting ranges to full values...")
	reducedFreshRanges := reduceFreshRanges(freshRanges)

	fmt.Println("Calculating answer part one...")
	totalIds := countFreshProducts(reducedFreshRanges, productIds)
	fmt.Println("Total fresh products:", totalIds)

	fmt.Println("Calculating answer part two...")
	totalFreshProducts := countTotalFreshProducts(reducedFreshRanges)
	fmt.Println("Total fresh products:", totalFreshProducts)

	executionTime := time.Since(start)
	fmt.Println("Execution time:", executionTime)
}

func readFreshRangesAndProducts() (freshRanges []Range, products []int) {
	readFile, err := os.Open("./real_input.txt")
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	passedEmptyLine := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		if line == "" {
			passedEmptyLine = true
			continue
		}

		if !passedEmptyLine {
			splitString := strings.Split(line, "-")
			start, _ := strconv.Atoi(splitString[0])
			end, _ := strconv.Atoi(splitString[1])
			freshRange := Range{start, end}

			freshRanges = append(freshRanges, freshRange)
		} else {
			value, _ := strconv.Atoi(line)
			products = append(products, value)
		}
	}

	readFile.Close()

	return freshRanges, products
}

func reduceFreshRanges(freshRanges []Range) []Range {
	if len(freshRanges) <= 1 {
		return freshRanges
	}

	// Sort ranges by Start
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].start < freshRanges[j].start
	})

	merged := []Range{freshRanges[0]}
	for i := 1; i < len(freshRanges); i++ {
		last := &merged[len(merged)-1]
		current := freshRanges[i]
		if current.start <= last.end {
			// Overlapping or adjacent, merge them
			last.end = max(last.end, current.end)
		} else {
			merged = append(merged, current)
		}
	}
	return merged
}

func countFreshProducts(freshRanges []Range, products []int) (freshProductsCount int) {
	for _, product := range products {
		for _, freshRange := range freshRanges {
			if product >= freshRange.start && product <= freshRange.end {
				// Product is fresh
				freshProductsCount++
			}
		}
	}

	return freshProductsCount
}

func countTotalFreshProducts(freshRanges []Range) (totalFreshProducts int) {
	for _, product := range freshRanges {
		totalFreshProducts += product.end - product.start + 1
	}

	return totalFreshProducts
}
