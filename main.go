package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day13"
)

func main() {
	s := common.WithTiming(day13.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
