package day17

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type C [4]int

func (c C) Add(q C) C {
	var res C
	for i, v := range c {
		res[i] = v + q[i]
	}
	return res
}

type Grid map[C]bool

func (g Grid) NumNeighbours(c C, del []C) int {
	var n int

	for _, d := range del {
		if g[c.Add(d)] {
			n++
		}
	}

	return n
}

func (g Grid) Step(del []C) Grid {
	// expand
	for p := range g {
		for _, d := range del {
			g[p.Add(d)] = g[p.Add(d)]
		}
	}

	next := make(Grid)
	for c, r := range g {
		numNeighbours := g.NumNeighbours(c, del)
		if r {
			next[c] = numNeighbours == 2 || numNeighbours == 3
		} else {
			next[c] = numNeighbours == 3
		}
	}
	return next
}

func delta(dim int) []C {
	if dim == 0 {
		return []C{{}}
	}
	var res []C
	for _, v := range []int{0, 1, -1} {
		for _, p := range delta(dim - 1) {
			p[dim-1] = v
			res = append(res, p)
		}
	}
	return res
}

type Solver struct{}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 17"
}

func SolvePartOne() error {
	grid, err := _read()
	if err != nil {
		return err
	}

	del := delta(3)[1:]

	for c := 0; c < 6; c++ {
		grid = grid.Step(del)
	}

	var acc int
	for _, r := range grid {
		if r {
			acc++
		}
	}

	fmt.Println(acc)

	return nil
}

func SolvePartTwo() error {
	grid, err := _read()
	if err != nil {
		return err
	}

	del := delta(4)[1:]

	for c := 0; c < 6; c++ {
		grid = grid.Step(del)
	}

	var acc int
	for _, r := range grid {
		if r {
			acc++
		}
	}

	fmt.Println(acc)

	return nil
}

func _read() (Grid, error) {
	i, err := ioutil.ReadFile("data/17.txt")
	if err != nil {
		return nil, err
	}

	fields := strings.Fields(string(i))

	grid := make(Grid)

	for y, s := range fields {
		for x, r := range s {
			grid[C{x, y}] = r == '#'
		}
	}
	return grid, nil
}
