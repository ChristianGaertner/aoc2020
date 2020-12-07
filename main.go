package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/day07"
)

func main() {
	s := common.WithTiming(day07.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
