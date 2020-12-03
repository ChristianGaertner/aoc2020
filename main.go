package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/daythree"
)

func main() {
	s := common.WithTiming(daythree.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
