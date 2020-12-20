package day20

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"os"
	"strconv"
)

type Tile struct {
	ID   uint64
	Data [][]bool

	Neighbours map[image.Point]*Tile
}

func NewTile(id uint64) *Tile {
	return &Tile{
		ID:         id,
		Neighbours: make(map[image.Point]*Tile),
	}
}

func (t *Tile) String() string {
	var res string
	for _, r := range t.Data {
		for _, c := range r {
			if c {
				res += "#"
			} else {
				res += "."
			}
		}
		res += "\n"
	}
	return res
}

func (t *Tile) Flip() {
	t.Data = Flip(t.Data)
}

func (t *Tile) Rotate() {
	t.Data = Rotate(t.Data)
}

func (t *Tile) HasRightNeighbour(o *Tile) bool {
	size := len(t.Data)

	for y := range t.Data {
		if t.Data[y][size-1] != o.Data[y][0] {
			return false
		}
	}

	return true
}

func (t *Tile) HasBottomNeighbour(o *Tile) bool {
	size := len(t.Data)

	for x := range t.Data[size-1] {
		if t.Data[size-1][x] != o.Data[0][x] {
			return false
		}
	}

	return true
}

var (
	N = image.Pt(0, -1)
	E = image.Pt(1, 0)
	S = image.Pt(0, 1)
	W = image.Pt(-0, 0)
)

func (t *Tile) HasNeighbor(o *Tile, modify bool) bool {
	for _, c := range []image.Point{N, E, S, W} {
		if _, ok := t.Neighbours[c]; ok {
			continue
		}

		for i := 0; i < 8; i++ {
			if c == E && t.HasRightNeighbour(o) {
				t.Neighbours[E] = o
				o.Neighbours[W] = t
				return true
			}
			if c == W && o.HasRightNeighbour(t) {
				t.Neighbours[W] = o
				o.Neighbours[E] = t
				return true
			}
			if c == S && t.HasBottomNeighbour(o) {
				t.Neighbours[S] = o
				o.Neighbours[N] = t
				return true
			}
			if c == N && o.HasBottomNeighbour(t) {
				t.Neighbours[N] = o
				o.Neighbours[S] = t
				return true
			}

			if !modify {
				break
			}

			if i%2 == 0 {
				o.Flip()
			} else {
				o.Flip()
				o.Rotate()
			}
		}
	}

	return false
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
	return "2020 12 20"
}

func SolvePartOne() error {
	allTiles, err := _read()
	if err != nil {
		return err
	}

	tm := make(map[uint64]*Tile)
	for _, t := range allTiles {
		tm[t.ID] = t
	}

	aligned := make(map[uint64]bool)
	for id, tile := range tm {
		if len(tile.Neighbours) != 0 {
			aligned[id] = true
		}
	}

	if len(aligned) == 0 {
		for tile := range tm {
			aligned[tile] = true
			break
		}
	}

	found := true
	for found {
		found = false
		for tile := range aligned {
			for other := range tm {
				if other == tile {
					continue
				}

				_, l := aligned[other]

				if tm[tile].HasNeighbor(tm[other], !l) {
					found = true
					aligned[other] = true
				}
			}
		}
	}

	var acc uint64
	acc += 1
	for _, t := range tm {
		if len(t.Neighbours) == 2 {
			acc *= t.ID
		}
	}
	fmt.Println(acc)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]*Tile, error) {
	file, err := os.Open("data/20.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tiles []*Tile

	var current *Tile

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}
		if raw[0] == ';' {
			continue
		}

		if raw[0] == 'T' {
			if current != nil {
				tiles = append(tiles, current)
			}

			n, err := strconv.Atoi(raw[5 : len(raw)-1])
			if err != nil {
				return nil, err
			}

			current = NewTile(uint64(n))
			continue
		}
		if current == nil {
			return nil, errors.New("data without tile")
		}

		var row []bool
		for _, c := range raw {
			row = append(row, c == '#')
		}

		current.Data = append(current.Data, row)
	}
	if current != nil {
		tiles = append(tiles, current)
	}

	return tiles, nil
}

func Rotate(data [][]bool) [][]bool {
	size := len(data)
	rotated := make([][]bool, size)

	for y := 0; y < size; y++ {
		rotated[y] = make([]bool, size)
	}

	for y := range data {
		for x := range data[y] {
			rotated[size-x-1][y] = data[y][x]
		}
	}

	return rotated
}

func Flip(data [][]bool) [][]bool {
	size := len(data)
	flipped := make([][]bool, size)

	for i := 0; i < size; i++ {
		flipped[i] = make([]bool, size)
	}

	for y := range data {
		for x := range data[y] {
			flipped[y][size-x-1] = data[y][x]
		}
	}

	return flipped
}
