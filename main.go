package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day03"
)

func main() {
	s := common.WithTiming(day03.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
