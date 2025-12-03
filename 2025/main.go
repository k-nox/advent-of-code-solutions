// Code generated; DO NOT EDIT.
// This file was generated at
// 2025-12-02 16:15:08.490708768 -0800 PST m=+0.001092529
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2025/day01"
	"github.com/k-nox/advent-of-code-solutions/2025/day02"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
	"day02": {PartOne: day02.PartOne, PartTwo: day02.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
