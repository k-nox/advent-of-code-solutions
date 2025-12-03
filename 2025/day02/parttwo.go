package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func PartTwo(useSample bool) int {
	idRanges := read(useSample)

	ivalidIDSum := 0

	for _, idRange := range idRanges {
		startRaw, endRaw, ok := strings.Cut(idRange, "-")
		if !ok {
			fmt.Printf("something weird happened trying to read line: %s\n", idRange)
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

		ivalidIDSum += sumInvalidIdsInRangePartTwo(start, end)
	}

	return ivalidIDSum
}

// maybe this would break down with larger digits, but it seems perfectly fine with 10 or less digits.
func sumInvalidIdsInRangePartTwo(start int, end int) int {
	count := 0

	for id := start; id <= end; id++ {
		if id < 11 {
			continue
		}
		idStr := strconv.Itoa(id)

		// if divisible by 2, check midway
		if len(idStr)%2 == 0 {
			midway := len(idStr) / 2
			if idStr[:midway] == idStr[midway:] {
				count += id
				continue
			}
		}

		// if divisible by 3, check thirds
		if len(idStr)%3 == 0 {
			firstThirdIdx := len(idStr) / 3
			if strings.Repeat(idStr[:firstThirdIdx], 3) == idStr {
				count += id
				continue
			}
		}
		// if divisible by 5, check fifths
		if len(idStr)%5 == 0 {
			firstFifthIdx := len(idStr) / 5
			if strings.Repeat(idStr[:firstFifthIdx], 5) == idStr {
				count += id
				continue
			}
		}
		// else can only be valid if all digits are the same (at this point, should only be strings of length 7, assuming a max length of 10)
		if strings.Repeat(idStr[:1], len(idStr)) == idStr {
			count += id
		}
	}

	return count
}
