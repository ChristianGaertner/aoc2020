package day12

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"strconv"
)

type Command struct {
	Action Direction
	Amount int
}

func (c Command) Vector() image.Point {
	if c.Action == North {
		return image.Pt(c.Amount, 0)
	}
	if c.Action == East {
		return image.Pt(0, c.Amount)
	}
	if c.Action == South {
		return image.Pt(-c.Amount, 0)
	}
	if c.Action == West {
		return image.Pt(0, -c.Amount)
	}
	return image.Pt(0, 0)
}

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"

	Left    Direction = "L"
	Right   Direction = "R"
	Forward Direction = "F"
)

var compass = [...]Direction{North, East, South, West}

type Ship struct {
	Dir Direction
	Pos image.Point
}

func (s Ship) ApplyCommand(cmd Command) Ship {
	if v := cmd.Vector(); !v.Eq(image.Pt(0, 0)) {
		return Ship{
			Dir: s.Dir,
			Pos: s.Pos.Add(v),
		}
	}
	if cmd.Action == Forward {
		return Ship{
			Dir: s.Dir,
			Pos: s.Pos.Add(Command{Action: s.Dir, Amount: cmd.Amount}.Vector()),
		}
	}

	turns := cmd.Amount / 90

	if cmd.Action == Left {
		turns = -turns
	}

	var index int
	for i, c := range compass {
		if s.Dir == c {
			index = i
		}
	}

	return Ship{
		Dir: compass[modLikePython(index+turns, 4)],
		Pos: s.Pos,
	}
}

func modLikePython(d, m int) int {
	var res = d % m
	if (res < 0 && m > 0) || (res > 0 && m < 0) {
		return res + m
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
	return "2020 12 12"
}

func SolvePartOne() error {
	cmds, err := _read()
	if err != nil {
		return err
	}

	ship := Ship{
		Dir: East,
	}

	for _, cmd := range cmds {
		ship = ship.ApplyCommand(cmd)
	}

	x := ship.Pos.X
	if x < 0 {
		x = -x
	}

	y := ship.Pos.Y
	if y < 0 {
		y = -y
	}

	fmt.Println(x + y)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]Command, error) {
	file, err := os.Open("data/12.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows []Command

	for scanner.Scan() {
		raw := scanner.Text()

		s := raw[0]
		n, err := strconv.Atoi(raw[1:])
		if err != nil {
			return nil, err
		}

		rows = append(rows, Command{
			Action: Direction(s),
			Amount: n,
		})
	}

	return rows, nil
}
