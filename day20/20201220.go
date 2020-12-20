package day20

import (
	"bufio"
	"errors"
	"fmt"
	"image"
	"os"
	"strconv"
)

type Image [][]bool

func NewImage(size image.Point, v bool) Image {
	var img Image
	for y := 0;  y < size.Y; y++ {
		var row []bool
		for x := 0; x < size.X; x++ {
			row = append(row, v)
		}
		img = append(img, row)
	}
	return img
}

func (t Image) Get(p image.Point) bool {
	if p.Y >= len(t) {
		return false
	}
	if p.X >= len(t[p.Y]) {
		return false
	}

	return t[p.Y][p.X]
}
func (t Image) Set(p image.Point, v bool) {
	t[p.Y][p.X] = v
}

func (t Image) String() string {
	var res string
	for _, r := range t {
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

type Tile struct {
	ID   uint64
	Data Image

	Neighbours map[image.Point]*Tile
}

func NewTile(id uint64) *Tile {
	return &Tile{
		ID:         id,
		Neighbours: make(map[image.Point]*Tile),
	}
}

func (t *Tile) Flip() {
	t.Data = Flip(t.Data)
}

func (t *Tile) Rotate() {
	t.Data = Rotate(t.Data)
}

func (t *Tile) HasRightNeighbour(o *Tile) bool {
	for y := range t.Data {
		if t.Data.Get(image.Pt(len(t.Data)-1, y)) != o.Data.Get(image.Pt(0, y)) {
			return false
		}
	}

	return true
}

func (t *Tile) HasBottomNeighbour(o *Tile) bool {
	for x := range t.Data[len(t.Data)-1] {
		if t.Data.Get(image.Pt(x, len(t.Data)-1)) != o.Data.Get(image.Pt(x, 0)) {
			return false
		}
	}

	return true
}

const numPermutations = 8

var (
	N = image.Pt(0, -1)
	E = image.Pt(1, 0)
	S = image.Pt(0, 1)
	W = image.Pt(-1, 0)
)

func (t *Tile) HasNeighbor(o *Tile, modify bool) bool {
	for _, dir := range []image.Point{N, E, S, W} {
		if _, ok := t.Neighbours[dir]; ok {
			continue
		}

		for i := 0; i < numPermutations; i++ {
			if dir == N && o.HasBottomNeighbour(t) {
				t.Neighbours[N] = o
				o.Neighbours[S] = t
				return true
			}
			if dir == E && t.HasRightNeighbour(o) {
				t.Neighbours[E] = o
				o.Neighbours[W] = t
				return true
			}
			if dir == S && t.HasBottomNeighbour(o) {
				t.Neighbours[S] = o
				o.Neighbours[N] = t
				return true
			}
			if dir == W && o.HasRightNeighbour(t) {
				t.Neighbours[W] = o
				o.Neighbours[E] = t
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

func alignTiles(tiles []*Tile) map[uint64]*Tile {
	tm := make(map[uint64]*Tile)
	for _, t := range tiles {
		tm[t.ID] = t
	}

	locked := make(map[uint64]bool)
	// align a random tile
	for tile := range tm {
		locked[tile] = true
		break
	}

	found := true
	for found {
		found = false
		for tile := range locked {
			for other := range tm {
				if other == tile {
					continue
				}

				_, l := locked[other]

				if tm[tile].HasNeighbor(tm[other], !l) {
					found = true
					locked[other] = true
				}
			}
		}
	}
	return tm
}

func Assemble(tiles map[uint64]*Tile) Image {
	visited := make(map[uint64]bool)
	q := make(map[uint64]image.Point)

	var startTile *Tile

	// visit the first tile
	for _, t := range tiles {
		startTile = t
		visited[startTile.ID] = true
		for n, p := range startTile.Neighbours {
			q[p.ID] = n
		}
		break
	}

	if startTile == nil {
		panic("WTF")
	}

	grid := map[image.Point]*Tile{
		{}: startTile,
	}

	var minX, minY, maxX, maxY int
	ts := len(startTile.Data) - 2
	for len(q) > 0 {
		for id, dir := range q {
			delete(q, id)
			visited[id] = true

			for n, np := range tiles[id].Neighbours {
				if _, ok := visited[np.ID]; !ok {
					q[np.ID] = dir.Add(n)
				}
			}

			grid[dir] = tiles[id]

			if dir.X < minX {
				minX = dir.X
			}
			if dir.Y < minY {
				minY = dir.Y
			}
			if dir.X > maxX {
				maxX = dir.X
			}
			if dir.Y > maxY {
				maxY = dir.Y
			}
		}
	}

	borderWidth := 1
	gs := maxX - minX + 1
	dim := image.Pt(gs * ts, (maxY-minY+1)*ts)
	img := NewImage(dim, false)
	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			for tY, row := range grid[image.Pt(x, y)].Data[borderWidth : ts+1] {
				dY := (y-minY)*ts + tY
				for tX, v := range row[borderWidth : len(row)-1] {
					dX := (x-minX)*ts + tX
					img.Set(image.Pt(dX, dY), v)
				}
			}
		}
	}
	return img
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
	tiles, err := _read()
	if err != nil {
		return err
	}

	aligned := alignTiles(tiles)

	var acc uint64
	acc += 1
	for _, t := range aligned {
		if len(t.Neighbours) == 2 {
			acc *= t.ID
		}
	}
	fmt.Println(acc)

	return nil
}

func SolvePartTwo() error {
	tiles, err := _read()
	if err != nil {
		return err
	}

	aligned := alignTiles(tiles)
	img := Assemble(aligned)
	seaMonster := make(map[image.Point]bool)

	for y, l := range []string{"                  # ", "#    ##    ##    ###", " #  #  #  #  #  #   "} {
		for x, c := range l {
			seaMonster[image.Pt(x, y)] = c == '#'
		}
	}

	for i := 0; i < numPermutations; i++ {
		found := make(map[image.Point]bool)
		for y := range img {
			for x := range img[y] {
				p := image.Pt(x, y)
				f := true
				for m, M := range seaMonster {
					if !M {
						continue
					}
					if !img.Get(p.Add(m)) {
						f = false
						break
					}

				}
				if f {
					found[p] = true
					// black out image part
					for m, M := range seaMonster {
						img.Set(p.Add(m), !(M || !img.Get(p.Add(m))))
					}
				}
			}
		}
		if len(found) > 0 {
			break
		}
		if i%2 == 0 {
			img = Flip(img)
		} else {
			img = Flip(img)
			img = Rotate(img)
		}
	}

	var acc uint64

	for _, yy := range img {
		for _, xx := range yy {
			if xx {
				acc++
			}
		}
	}

	fmt.Println(acc)

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

func Rotate(data Image) Image {
	size := len(data)
	rotated := make(Image, size)

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

func Flip(data Image) Image {
	size := len(data)
	flipped := make(Image, size)

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
