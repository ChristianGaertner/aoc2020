package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/daytwo"
)

func main() {
	s := common.WithTiming(daytwo.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}
