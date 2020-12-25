package day24

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strings"
)

type Solver struct{}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 24"
}

func Neighbours() [6]image.Point {
	return [6]image.Point{
		image.Pt(2, 0),
		image.Pt(1, 1),
		image.Pt(-1, 1),
		image.Pt(-2, 0),
		image.Pt(-1, -1),
		image.Pt(1, -1),
	}
}

func resolve(ins string, ref image.Point) (string, image.Point) {
	if strings.HasPrefix(ins, "e") {
		return resolve(ins[1:], ref.Add(image.Pt(2, 0)))
	}
	if strings.HasPrefix(ins, "se") {
		return resolve(ins[2:], ref.Add(image.Pt(1, 1)))
	}
	if strings.HasPrefix(ins, "sw") {
		return resolve(ins[2:], ref.Add(image.Pt(-1, 1)))
	}
	if strings.HasPrefix(ins, "w") {
		return resolve(ins[1:], ref.Add(image.Pt(-2, 0)))
	}
	if strings.HasPrefix(ins, "nw") {
		return resolve(ins[2:], ref.Add(image.Pt(-1, -1)))
	}
	if strings.HasPrefix(ins, "ne") {
		return resolve(ins[2:], ref.Add(image.Pt(1, -1)))
	}
	return ins, ref
}

func SolvePartOne() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	// true == black
	tiles := make(map[image.Point]bool)

	for _, ins := range instructions {
		_, target := resolve(ins, image.Pt(0, 0))

		_, ok := tiles[target]
		if ok {
			delete(tiles, target)
		} else {
			tiles[target] = true
		}
	}

	var acc int
	for _, isBlack := range tiles {
		if isBlack {
			acc++
		}
	}

	fmt.Println(acc)
	return nil
}

func SolvePartTwo() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	// true == black
	tiles := make(map[image.Point]bool)

	for _, ins := range instructions {
		_, target := resolve(ins, image.Pt(0, 0))

		_, ok := tiles[target]
		if ok {
			delete(tiles, target)
		} else {
			tiles[target] = true
		}
	}

	for i := 0; i < 100; i++ {
		neigh := map[image.Point]int{}
		for ref := range tiles {
			for _, n := range Neighbours() {
				target := ref.Add(n)
				neigh[target]++
			}
		}

		next := make(map[image.Point]bool)
		for ref, n := range neigh {
			if _, ok := tiles[ref]; ok && n == 1 || n == 2 {
				next[ref] = true
			}
		}
		tiles = next
	}

	fmt.Println(len(tiles))

	return nil
}

func _read() ([]string, error) {
	file, err := os.Open("data/24.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var res []string

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}
		if raw[0] == ';' {
			continue
		}

		res = append(res, raw)
	}
	return res, nil
}
