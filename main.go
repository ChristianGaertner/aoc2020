package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day14"
)

func main() {
	s := common.WithTiming(day14.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
