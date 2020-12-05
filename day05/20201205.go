package day05

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type BSP string

func (b BSP) RowPart() string {
	return string(b[:7])
}

func (b BSP) SeatPart() string {
	return string(b[7:])
}

func (b BSP) RowAndSeat() (int, int) {
	lower := 0
	upper := 127

	for _, c := range b.RowPart() {
		if c == 'F' {
			upper = int(math.Floor(float64(lower+upper) / 2.0))
		} else {
			lower = int(math.Round(float64(lower+upper) / 2.0))
		}
	}
	if lower != upper {
		panic("MISMATCH ROW FOR " + b)
	}

	left := 0
	right := 7
	for _, c := range b.SeatPart() {
		if c == 'L' {
			right = int(math.Floor(float64(left+right) / 2.0))
		} else {
			left = int(math.Round(float64(left+right) / 2.0))
		}
	}

	if left != right {
		panic("MISMATCH SEAT FOR " + b)
	}

	return lower, left
}

func (b BSP) ID() int {
	row, seat := b.RowAndSeat()

	return row*8 + seat
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
	return "2020 12 05"
}

func SolvePartOne() error {
	bsps, err := _read()
	if err != nil {
		return err
	}

	var high int

	for _, bsp := range bsps {
		id := bsp.ID()
		if id > high {
			high = id
		}
	}

	fmt.Printf("high=%d\n", high)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]BSP, error) {
	file, err := os.Open("data/05.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var bsps []BSP

	for scanner.Scan() {
		raw := scanner.Text()
		bsps = append(bsps, BSP(raw))
	}

	return bsps, nil
}
