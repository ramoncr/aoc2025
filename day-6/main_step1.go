package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("Advent of Code - Day 6")
	fmt.Println("=========================")
	fmt.Println("Reading input contents...")
	exerciseGrid := readFileContentsAndSplit()
	exerciseGrid = cleanSpaces(exerciseGrid)
	_, totalAnswer := calculateAnswers(exerciseGrid)

	fmt.Println("Total value:", totalAnswer)

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}

func readFileContentsAndSplit() [][]string {
	readFile, err := os.Open("real_input.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var instructions [][]string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		instructions = append(instructions, strings.Split(line, " "))
	}

	return instructions
}

func cleanSpaces(input [][]string) [][]string {
	var result [][]string

	for _, line := range input {
		result = append(result, removeAll(line, " "))
	}

	return result
}

func removeAll(s []string, val string) []string {
	var result []string
	for _, v := range s {
		if strings.Trim(v, " ") != "" {
			result = append(result, v)
		}
	}
	return result
}

func calculateAnswers(instructions [][]string) ([]int, int) {
	// Determine amount of calculations
	amountOfCalculations := len(instructions[0])
	amountOfInputs := len(instructions) - 1 // last one is instruction
	var values []int
	var totalValue int

	for i := 0; i < amountOfCalculations; i++ {
		var instructionChar = instructions[len(instructions)-1][i]
		operator := operationsMap[instructionChar]

		calculationValue := 0
		if instructionChar == "*" {
			calculationValue = 1
		}

		for j := 0; j < amountOfInputs; j++ {
			currentNumber, _ := strconv.Atoi(instructions[j][i])
			calculationValue = operator(calculationValue, currentNumber)
		}

		values = append(values, calculationValue)
		totalValue += calculationValue
	}

	return values, totalValue
}
