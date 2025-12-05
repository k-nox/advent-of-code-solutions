package day04

import (
	"bufio"
	"fmt"
	"image"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	f := helper.OpenInput(2025, 4, useSample)
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	scanner := bufio.NewScanner(f)
	grid := parseGrid(scanner)

	rolls := grid.Members()
	directionsToCheck := []helper.Direction{helper.Left, helper.Right, helper.Up, helper.Down, helper.UpLeft, helper.DownLeft, helper.UpRight, helper.DownRight}
	canBeReached := 0
	for _, roll := range rolls {
		otherRollsCount := 0
		for _, direction := range directionsToCheck {
			if otherRollsCount == 4 {
				break
			}
			if grid.Contains(direction(roll)) {
				otherRollsCount++
			}
		}
		if otherRollsCount < 4 {
			canBeReached++
		}
	}

	return canBeReached
}

func parseGrid(scanner *bufio.Scanner) helper.Set[image.Point] {
	grid := helper.Set[image.Point]{}
	row := 0
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		for col, char := range line {
			if char == '@' {
				grid.Add(image.Point{X: row, Y: col})
			}
		}
		row++
	}
	return grid
}
