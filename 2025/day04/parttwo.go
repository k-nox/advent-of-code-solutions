package day04

import (
	"bufio"
	"fmt"
	"image"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2025, 4, useSample)
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	scanner := bufio.NewScanner(f)

	grid := parseGrid(scanner)

	totalRemoved := 0
	for grid.Len() > 0 {
		var removedCount int
		grid, removedCount = removeRolls(grid)
		if removedCount == 0 {
			break
		}
		totalRemoved += removedCount
	}

	return totalRemoved
}

func removeRolls(grid helper.Set[image.Point]) (helper.Set[image.Point], int) {
	removedCount := 0
	rolls := grid.Members()
	for _, roll := range rolls {
		otherRollsCount := 0
		for _, direction := range helper.AllDirections {
			if otherRollsCount == 4 {
				break
			}
			if grid.Contains(direction(roll)) {
				otherRollsCount++
			}
		}
		if otherRollsCount < 4 {
			grid.Remove(roll)
			removedCount++
		}
	}
	return grid, removedCount
}
