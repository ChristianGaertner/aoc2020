package day03

import (
	"bufio"
	"fmt"
	"os"
)

type Solver struct{}

type Map struct {
	Rows []*Row
}

func (m *Map) String() string {
	var res string

	for _, row := range m.Rows {
		res += row.String() + "\n"
	}

	return res
}

func (m *Map) Height() int {
	return len(m.Rows)
}

func (m *Map) IsTreeAt(x, y int) bool {
	return m.Rows[y].IsTreeAt(x)
}

type Row struct {
	IsTree []bool
}

func (r *Row) IsTreeAt(y int) bool {
	return r.IsTree[y%len(r.IsTree)]
}

func (r *Row) String() string {
	var res string

	for _, isTree := range r.IsTree {
		if isTree {
			res += "#"
		} else {
			res += "."
		}
	}

	return res
}

type Slope struct {
	DX int
	DY int
}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 03"
}

func SolvePartOne() error {
	mp, err := _read()
	if err != nil {
		return err
	}

	slope := Slope{
		DX: 3,
		DY: 1,
	}

	var numTrees int

	x := 0

	for y := 0; y < mp.Height(); y += slope.DY {
		if mp.IsTreeAt(x, y) {
			numTrees += 1
		}

		x += slope.DX
	}

	fmt.Printf("numTrees=%d\n", numTrees)
	return nil
}

func SolvePartTwo() error {
	mp, err := _read()
	if err != nil {
		return err
	}

	slopes := []Slope{
		{DX: 1, DY: 1},
		{DX: 3, DY: 1},
		{DX: 5, DY: 1},
		{DX: 7, DY: 1},
		{DX: 1, DY: 2},
	}

	res := 1

	for _, slope := range slopes {
		var numTrees int

		x := 0

		for y := 0; y < mp.Height(); y += slope.DY {
			if mp.IsTreeAt(x, y) {
				numTrees += 1
			}

			x += slope.DX
		}

		res *= numTrees
	}

	fmt.Printf("result=%d\n", res)
	return nil
}

func _read() (*Map, error) {
	file, err := os.Open("data/03.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var rows []*Row

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		row := &Row{}
		for _, c := range raw {
			row.IsTree = append(row.IsTree, c == '#')
		}

		rows = append(rows, row)
	}

	return &Map{Rows: rows}, nil
}
