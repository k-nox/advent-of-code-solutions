package day03

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/k-nox/advent-of-code-solutions/helper"
)

func PartOne(useSample bool) int {
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
			totalJoltage += calculateJoltage([]rune(bankRaw))
		}
	}

	return totalJoltage
}

func calculateJoltage(bank []rune) int {
	maxBatteryIdx := 0

	if bank[maxBatteryIdx] != '9' {
		for idx := 1; idx < len(bank)-1; idx++ {
			if bank[idx] > bank[maxBatteryIdx] {
				maxBatteryIdx = idx
			}
		}
	}

	maxBatteryIdxFromTail := len(bank) - 1

	if bank[maxBatteryIdxFromTail] != '9' {
		for idx := len(bank) - 2; idx > maxBatteryIdx; idx-- {
			if bank[idx] > bank[maxBatteryIdxFromTail] {
				maxBatteryIdxFromTail = idx
			}
		}
	}

	joltageRaw := string([]rune{bank[maxBatteryIdx], bank[maxBatteryIdxFromTail]})
	joltage, err := strconv.Atoi(joltageRaw)
	if err != nil {
		fmt.Printf("something weird happened trying to convert joltage: %s: %s", joltageRaw, err.Error())
	}

	return joltage
}
