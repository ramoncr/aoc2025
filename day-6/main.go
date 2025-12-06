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

type Operation func(int, int) int

var operationsMap = map[string]Operation{
	"+": func(a int, b int) int { return a + b },
	"*": func(a int, b int) int { return a * b },
}

func main() {
	start := time.Now()
	fmt.Println("Advent of Code - Day 6")
	fmt.Println("=========================")
	fmt.Println("Reading input contents...")
	lines := readFileContents()

	calculations := groupColumns(lines)
	_, totalSum := calculate(calculations)

	fmt.Println("Total sum:", totalSum)

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}

func readFileContents() []string {
	readFile, err := os.Open("real_input.txt")
	defer readFile.Close()

	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var lines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func groupColumns(lines []string) (result [][]string) {
	line := lines[len(lines)-1]

	var operatorIndexes []int

	for i, char := range line {
		if char == '*' || char == '+' {
			operatorIndexes = append(operatorIndexes, i)
		}
	}

	for i := 0; i < len(operatorIndexes); i++ {
		var columnValues []string

		start := operatorIndexes[i]

		end := len(line)
		if i+1 < len(operatorIndexes) {
			end = operatorIndexes[i+1]
		}

		for j := 0; j < len(lines); j++ {
			number := lines[j][start:end]
			columnValues = append(columnValues, number)
		}

		result = append(result, columnValues)
	}

	return result
}

func calculate(calculations [][]string) (result []int, totalSum int) {
	totalSum = 0
	for _, calculationLines := range calculations {
		var numbers []int

		for charIndex := 0; charIndex < len(calculationLines[0]); charIndex++ {
			var rawNumber string

			// -1 is to avoid including the operator
			for lineIndex := 0; lineIndex < len(calculationLines)-1; lineIndex++ {
				rawNumber += string(calculationLines[lineIndex][charIndex])
			}
			parsedNumber, _ := strconv.Atoi(strings.Trim(rawNumber, " "))
			if parsedNumber != 0 {
				numbers = append(numbers, parsedNumber)
			}
		}

		operatorChar := strings.Trim(calculationLines[len(calculationLines)-1], " ")
		operator := operationsMap[operatorChar]

		lineResult := 0
		if operatorChar == "*" {
			lineResult = 1
		}

		for _, number := range numbers {
			lineResult = operator(lineResult, number)
		}

		result = append(result, lineResult)
		totalSum += lineResult

	}

	return result, totalSum
}
