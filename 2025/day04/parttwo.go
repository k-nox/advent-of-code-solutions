package day04

import (
	"bufio"
  "fmt"

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

	for scanner.Scan() {
	
	}

	return 0
}
