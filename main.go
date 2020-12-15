package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day15"
)

func main() {
	s := common.WithTiming(day15.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
