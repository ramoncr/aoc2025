package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Loading input...")
	grid := loadInput()

	fmt.Println("Counting paper roles inefficiently...")

	totalRemovedPackages := 0
	firstRun := true
	answerPart1 := 0
	for {
		accessiblePackages, updatedGrid := calculateAccessiblePackages(grid)
		grid = updatedGrid
		totalRemovedPackages += accessiblePackages

		if firstRun {
			answerPart1 = accessiblePackages
			firstRun = false
		}

		if accessiblePackages == 0 {
			break
		}
	}

	fmt.Println("Part 1, removed in the first run:", answerPart1)
	fmt.Println("Part 2, total removed after:", totalRemovedPackages)

	elapsed := time.Since(start)
	fmt.Println("Execution time:", elapsed)
}

func loadInput() [][]string {
	file, _ := os.Open("real_input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var lines [][]string
	for scanner.Scan() {
		line := scanner.Text()
		horizontalLine := strings.Split(line, "")
		lines = append(lines, horizontalLine)
	}

	return lines
}

func calculateAccessiblePackages(grid [][]string) (int, [][]string) {
	var accessiblePackages int

	var newgrid [][]string

	for rowIndex, row := range grid {

		var newRow []string
		for colIndex, rowValue := range row {
			if rowValue != "@" {
				//fmt.Print(rowValue)
				newRow = append(newRow, rowValue)
				continue
			}

			// Check all 8 positions
			rowAbove := rowIndex - 1
			rowCurrent := rowIndex
			rowBelow := rowIndex + 1
			columnLeft := colIndex - 1

			var surroundingPackages int

			// Check line above
			if rowAbove >= 0 {
				surroundingPackages += calculateRow(grid, rowAbove, columnLeft, false)
			}

			// check current row
			surroundingPackages += calculateRow(grid, rowCurrent, columnLeft, true)

			if rowBelow < len(grid) {
				// Row below exists count it
				surroundingPackages += calculateRow(grid, rowBelow, columnLeft, false)
			}

			if surroundingPackages < 4 {
				// Include in the count
				accessiblePackages++
				//fmt.Print("x")
				newRow = append(newRow, ".")
			} else {
				//fmt.Print("@")
				newRow = append(newRow, "@")
			}

		}

		//fmt.Println(" ")
		newgrid = append(newgrid, newRow)
	}

	return accessiblePackages, newgrid
}

func checkSpotForPackage(row []string, index int) bool {
	return index >= 0 && index < len(row) && row[index] == "@"
}

func calculateRow(grid [][]string, rowIndex int, leftColumIndex int, skipCentre bool) (counter int) {
	if rowIndex < 0 || rowIndex >= len(grid) {
		return counter
	}

	if checkSpotForPackage(grid[rowIndex], leftColumIndex) {
		counter++
	}
	if !skipCentre && checkSpotForPackage(grid[rowIndex], leftColumIndex+1) {
		counter++
	}

	if checkSpotForPackage(grid[rowIndex], leftColumIndex+2) {
		counter++
	}
	return counter
}
