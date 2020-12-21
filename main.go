package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day21"
)

func main() {
	s := common.WithTiming(day21.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
