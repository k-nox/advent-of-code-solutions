package day02

import (
	"bufio"
	"regexp"
	"slices"
	"strconv"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2015, 2, useSample)
	defer f.Close()

	scanner := bufio.NewScanner(f)

	re := regexp.MustCompile(`\d+`)
	total := 0
	for scanner.Scan() {
		dimensions := re.FindAllString(scanner.Text(), -1)
		l, err := strconv.Atoi(dimensions[0])
		if err != nil {
			panic(err)
		}
		w, err := strconv.Atoi(dimensions[1])
		if err != nil {
			panic(err)
		}
		h, err := strconv.Atoi(dimensions[2])
		if err != nil {
			panic(err)
		}
		total += slices.Min(perimeters(l, w, h))
		total += (l * w * h)
	}

	return total
}

func perimeters(l int, w int, h int) []int {
	return []int{2 * (l + w), 2 * (w + h), 2 * (h + l)}
}
