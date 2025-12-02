package day01

import (
	"bufio"
	"fmt"
	"math"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2025, 1, useSample)
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	scanner := bufio.NewScanner(f)

	dial := 50
	zeroCount := 0

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		count := decode(line)
		boundedCount := count % 100
		if boundedCount != count {
			zeroCount = zeroCount + (int(math.Abs(float64(count))) / 100)
		}
		prevDial := dial
		dial = dial + boundedCount
		if dial < 0 {
			dial = 100 + dial
			if prevDial != 0 {
				zeroCount++
			}
		} else if dial >= 100 {
			dial = dial - 100
			if prevDial != 0 {
				zeroCount++
			}
		} else if dial == 0 {
			zeroCount++
		}
	}

	return zeroCount
}
