package common

import (
	"fmt"
	"time"
)

type Solver interface {
	Solve() error
	Day() string
}

type withTiming struct {
	Solver Solver
}

func WithTiming(s Solver) Solver {
	return &withTiming{
		Solver: s,
	}
}

func (w *withTiming) Solve() error {
	startTime := time.Now()

	defer func() {
		duration := time.Now().Sub(startTime)
		fmt.Printf("---- END (%05dµs) ----\n", duration.Microseconds())
	}()
	fmt.Printf("----   %s  ----\n", w.Day())
	return w.Solver.Solve()
}

func (w *withTiming) Day() string {
	return w.Solver.Day()
}
