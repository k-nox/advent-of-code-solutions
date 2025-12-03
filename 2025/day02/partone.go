package day02

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
	idRanges := read(useSample)

	invalidIDSum := 0
	for _, idRange := range idRanges {
		startRaw, endRaw, ok := strings.Cut(idRange, "-")

		if !ok {
			fmt.Printf("something weird happened trying to read line: %s\n", idRange)
			break
		}

		if len(startRaw)%2 != 0 && len(startRaw) == len(endRaw) {
			// range only contains odd number of digits, so invalid ids are impossible
			continue
		}

		start, err := strconv.Atoi(strings.TrimSpace(startRaw))
		if err != nil {
			fmt.Printf("something weird happened trying to convert start range: %s: %s\n", startRaw, err.Error())
			break
		}
		end, err := strconv.Atoi(strings.TrimSpace(endRaw))
		if err != nil {
			fmt.Printf("something weird happened trying to convert end range: %s: %s\n", endRaw, err.Error())
			break
		}

		invalidIDSum += sumInvalidIdsInRange(start, end)

	}

	return invalidIDSum
}

func read(useSample bool) []string {
	return strings.Split(strings.TrimSpace(string(helper.ReadInput(2025, 2, useSample))), ",")
}

func sumInvalidIdsInRange(start int, end int) int {
	count := 0

	for id := start; id <= end; id++ {
		idStr := strconv.Itoa(id)
		if len(idStr)%2 != 0 {
			// odd number of digits, so can't be a repeat
			continue
		}

		midway := len(idStr) / 2

		if idStr[:midway] == idStr[midway:] {
			count += id
		}
	}
	return count
}
