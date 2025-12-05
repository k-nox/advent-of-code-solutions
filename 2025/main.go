// Code generated; DO NOT EDIT.
// This file was generated at
// 2025-12-04 17:41:56.887957536 -0800 PST m=+0.001324735
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2025/day01"
	"github.com/k-nox/advent-of-code-solutions/2025/day02"
	"github.com/k-nox/advent-of-code-solutions/2025/day03"
	"github.com/k-nox/advent-of-code-solutions/2025/day04"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	"day02": {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	"day03": {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
	"day04": {PartOne: day04.PartOne, PartTwo: day04.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
