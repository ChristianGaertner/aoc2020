package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day08"
)

func main() {
	s := common.WithTiming(day08.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
