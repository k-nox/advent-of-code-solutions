package day01

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
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
		dial = rotate(dial, count)
		if dial == 0 {
			zeroCount++
		}
	}

	return zeroCount
}

func rotate(from int, count int) int {
	boundedCount := count % 100
	result := from + boundedCount
	if result < 0 {
		result = 100 + result
	} else if result >= 100 {
		result = result - 100
	}

	return result
}

func decode(line string) int {
	if len(line) == 0 {
		return 0
	}
	num, err := strconv.Atoi(line[1:])
	if err != nil {
		panic(fmt.Errorf("line not convertible to number: %s: %w", line, err))
	}
	if line[0] == 'L' {
		return num * -1
	}
	return num
}
