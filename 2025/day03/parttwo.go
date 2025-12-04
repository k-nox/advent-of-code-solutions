package day03

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartTwo(useSample bool) int {
	f := helper.OpenInput(2025, 3, useSample)
	defer func() {
		if err := f.Close(); err != nil {
			panic(fmt.Errorf("failed to close file: %w", err))
		}
	}()

	scanner := bufio.NewScanner(f)

	totalJoltage := 0
	for scanner.Scan() {
		bankRaw := strings.TrimSpace(scanner.Text())
		if len(bankRaw) > 0 {
			totalJoltage += calculateJoltagePartTwo([]rune(bankRaw))
		}
	}

	return totalJoltage
}

func calculateJoltagePartTwo(bank []rune) int {
	joltageRaw := []rune{}
	lastBatteryIdx := -1
	for len(joltageRaw) < 12 {
		nextBatteryIdx := findNextBattery(bank, lastBatteryIdx, 12-len(joltageRaw))
		nextBattery := bank[nextBatteryIdx]
		joltageRaw = append(joltageRaw, nextBattery)
		lastBatteryIdx = nextBatteryIdx
	}

	joltage, err := strconv.Atoi(string(joltageRaw))
	if err != nil {
		fmt.Printf("something weird happened trying to convert joltage: %s: %s\n", string(joltageRaw), err)
	}
	return joltage
}

func findNextBattery(bank []rune, lastBatteryIdx int, maxLen int) int {
	maxBatteryIdx := lastBatteryIdx + 1
	if bank[maxBatteryIdx] != '9' {
		for idx := maxBatteryIdx + 1; idx <= len(bank)-maxLen; idx++ {
			if bank[idx] > bank[maxBatteryIdx] {
				maxBatteryIdx = idx
			}
		}
	}
	return maxBatteryIdx
}
