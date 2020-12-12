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

func (c Command) String() string {
	return fmt.Sprintf("%s%d", c.Action, c.Amount)
}

func (c Command) Vector() image.Point {
	if c.Action == North {
		return image.Pt(0, c.Amount)
	}
	if c.Action == East {
		return image.Pt(c.Amount, 0)
	}
	if c.Action == South {
		return image.Pt(0, -c.Amount)
	}
	if c.Action == West {
		return image.Pt(-c.Amount, 0)
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
	cmds, err := _read()
	if err != nil {
		return err
	}

	ship := image.Pt(0, 0)
	wp := image.Pt(10, -1)
	for _, cmd := range cmds {
		switch cmd.Action {
		case North:
			wp.Y -= cmd.Amount
		case South:
			wp.Y += cmd.Amount
		case East:
			wp.X += cmd.Amount
		case West:
			wp.X -= cmd.Amount
		case Left:
			for i := 0; i < cmd.Amount; i += 90 {
				wp.X, wp.Y = wp.Y, -wp.X
			}
		case Right:
			for i := 0; i < cmd.Amount; i += 90 {
				wp.X, wp.Y = -wp.Y, wp.X
			}
		case Forward:
			ship.X += wp.X * cmd.Amount
			ship.Y += wp.Y * cmd.Amount
		}
	}

	x := ship.X
	if x < 0 {
		x = -x
	}

	y := ship.Y
	if y < 0 {
		y = -y
	}

	fmt.Println(x + y)

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
