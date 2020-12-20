package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day20"
)

func main() {
	s := common.WithTiming(day20.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
