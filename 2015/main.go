// Code generated; DO NOT EDIT.
// This file was generated at
// 2024-12-12 14:11:43.370014 -0800 PST m=+0.540882084
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2015/day01"
	"github.com/k-nox/advent-of-code-solutions/2015/day02"
	"github.com/k-nox/advent-of-code-solutions/2015/day03"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	"day02": {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
	"day03": {PartOne: day03.PartOne, PartTwo: day03.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}