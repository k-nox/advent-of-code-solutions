// Code generated; DO NOT EDIT.
// This file was generated at
// 2025-12-01 16:59:29.323485138 -0800 PST m=+0.001101224
package main

import (
	"log"
	"os"
	"github.com/k-nox/aoc/cli"
	"github.com/k-nox/advent-of-code-solutions/2025/day01"
)

var registry = cli.Registry{
	"day01": {PartOne: day01.PartOne, PartTwo: day01.PartTwo},
}	

func main() {
	app := cli.App(registry)
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
