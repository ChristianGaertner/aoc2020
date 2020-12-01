package main

import (
	"github.com/ChristianGaertner/aoc2020/common"
	"github.com/ChristianGaertner/aoc2020/dayone"
)

func main() {
	s := common.WithTiming(dayone.Solver{})
	if err := s.Solve(); err != nil {
		panic(err)
	}
}