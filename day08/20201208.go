package day08

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Solver struct{}

type Instruction struct {
	Type  string
	Value int
}

type Interpreter struct {
	Acc int
}

func (i *Interpreter) Exec(ins []Instruction) (int, bool) {
	seen := make(map[int]bool)
	for p := 0; p < len(ins); {
		if seen[p] {
			return i.Acc, false
		}
		seen[p] = true
		c := ins[p]
		switch c.Type {
		case "nop":
			p++
			continue
		case "acc":
			i.Acc += c.Value
			p++
			continue
		case "jmp":
			p += c.Value
			continue
		}
	}
	return i.Acc, true
}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 08"
}

func SolvePartOne() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	i := Interpreter{}

	res, _ := i.Exec(instructions)

	fmt.Printf("acc=%d\n", res)

	return nil
}

func SolvePartTwo() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	var res int
	needsRecheck := true

	checkStart := len(instructions)

	tries := 0
	for needsRecheck && tries <= len(instructions) {
		tries++

		testSet := make([]Instruction, len(instructions))
		copy(testSet, instructions)

		for i := checkStart - 1; i >= 0; i-- {
			if testSet[i].Type == "jmp" {
				testSet[i].Type = "nop"
				checkStart = i
				break
			}
		}

		i := Interpreter{}
		r, terminated := i.Exec(testSet)
		if terminated {
			res = r
		}
		needsRecheck = !terminated
	}

	fmt.Printf("acc=%d\n", res)

	return nil
}

func _read() ([]Instruction, error) {
	file, err := os.Open("data/08.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []Instruction

	for scanner.Scan() {
		raw := scanner.Text()

		typ := raw[:3]
		v, err := strconv.Atoi(raw[4:])
		if err != nil {
			return nil, err
		}

		instructions = append(instructions, Instruction{
			Type:  typ,
			Value: v,
		})

	}

	return instructions, nil
}
