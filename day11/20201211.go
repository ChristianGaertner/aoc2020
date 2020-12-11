package day11

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

type FerrySeats []FerryRow

func (f FerrySeats) String() string {
	var res string

	for _, row := range f {
		res += row.String() + "\n"
	}

	return res
}

func (f FerrySeats) Eq(o FerrySeats) bool {
	return f.String() == o.String()
}

func (f FerrySeats) NumOccupied() int {
	var n int
	for _, r := range f {
		n += NumOccupied(r)
	}
	return n
}

func (f FerrySeats) Clone() FerrySeats {
	var next FerrySeats

	for _, row := range f {
		next = append(next, row.Clone())
	}

	return next
}

func (f FerrySeats) SetSeat(x, y int, v Seat) {
	f[y][x] = v
}

func (f FerrySeats) Adjacent(x, y int) []Seat {
	var ad []Seat

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dy == 0 && dx == 0 {
				continue
			}

			py := y + dy
			px := x + dx

			if py < 0 || py >= len(f) || px < 0 || px >= len(f[0]) {
				continue
			}

			ad = append(ad, f[py][px])
		}
	}
	return ad
}

func (f FerrySeats) AdjacentVisible(x, y int) []Seat {
	var ad []Seat

	for dx := -1; dx <= 1; dx++ {
	outer:
		for dy := -1; dy <= 1; dy++ {
			if dy == 0 && dx == 0 {
				continue
			}

			for i := 1; i < 1000; i++ {
				py := y + dy*i
				px := x + dx*i
				if py < 0 || py >= len(f) || px < 0 || px >= len(f[0]) {
					continue outer
				}

				s := f[py][px]
				if s == SeatFloor {
					continue
				}
				ad = append(ad, f[py][px])
				continue outer
			}
		}
	}
	return ad
}

func NumOccupied(seats []Seat) int {
	var n int
	for _, s := range seats {
		if s == SeatOccupied {
			n++
		}
	}
	return n
}

func NoOccupied(seats []Seat) bool {
	for _, s := range seats {
		if s == SeatOccupied {
			return false
		}
	}
	return true
}

type FerryRow []Seat

func (f FerryRow) String() string {
	var res string

	for _, s := range f {
		if s == SeatFloor {
			res += "."
		}
		if s == SeatEmpty {
			res += "L"
		}
		if s == SeatOccupied {
			res += "#"
		}
	}

	return res
}

func (f FerryRow) Clone() FerryRow {
	var next FerryRow

	for _, s := range f {
		next = append(next, s)
	}

	return next
}

type Seat int

const (
	SeatFloor Seat = iota + 1
	SeatEmpty
	SeatOccupied
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
	return "2020 12 11"
}

func SolvePartOne() error {
	ferry, err := _read()
	if err != nil {
		return err
	}

	prev := ferry
	for i := 0; i < 100; i++ {
		next := prev.Clone()
		for y, row := range prev {
			for x, column := range row {
				adj := prev.Adjacent(x, y)
				if column == SeatEmpty && NoOccupied(adj) {
					next.SetSeat(x, y, SeatOccupied)
				} else if column == SeatOccupied && NumOccupied(adj) >= 4 {
					next.SetSeat(x, y, SeatEmpty)
				}
			}
		}

		if prev.Eq(next) {
			break
		}
		prev = next
	}

	fmt.Println(prev.NumOccupied())

	return nil
}

func SolvePartTwo() error {
	ferry, err := _read()
	if err != nil {
		return err
	}

	prev := ferry
	for i := 0; i < 100; i++ {
		next := prev.Clone()

		for y, row := range prev {
			for x, column := range row {
				adj := prev.AdjacentVisible(x, y)
				if column == SeatEmpty && NoOccupied(adj) {
					next.SetSeat(x, y, SeatOccupied)
				} else if column == SeatOccupied && NumOccupied(adj) >= 5 {
					next.SetSeat(x, y, SeatEmpty)
				}
			}
		}

		if prev.Eq(next) {
			break
		}
		prev = next
	}

	fmt.Println(prev.NumOccupied())

	return nil
}

func _read() (FerrySeats, error) {
	file, err := os.Open("data/11.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows []FerryRow

	for scanner.Scan() {
		raw := scanner.Text()
		var seats FerryRow

		for _, r := range strings.Split(raw, "") {
			if r == "." {
				seats = append(seats, SeatFloor)
			} else if r == "L" {
				seats = append(seats, SeatEmpty)
			} else if r == "#" {
				seats = append(seats, SeatOccupied)
			} else {
				return nil, errors.New("unknown seat type")
			}
		}

		rows = append(rows, seats)
	}

	return rows, nil
}
