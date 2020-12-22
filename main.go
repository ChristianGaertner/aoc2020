package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day22"
)

func main() {
	s := common.WithTiming(day22.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
