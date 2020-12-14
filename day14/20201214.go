package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mask []byte

type Instruction struct {
	Mask Mask
	Set  [2]int
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
	return "2020 12 14"
}

func SolvePartOne() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	register := make(map[int]int)

	mask := 0
	mmask := 0

	for _, ins := range instructions {
		if len(ins.Mask) != 0 {
			mask = 0
			mmask = 0
			for _, c := range ins.Mask {
				mask = mask << 1
				mmask = mmask << 1
				if c != 'X' {
					mmask |= 1
					if c == '1' {
						mask |= 1
					}
				}
			}
			continue
		}
		addr := ins.Set[0]
		v := ins.Set[1]
		register[addr] = (v &^ mmask) | (mask & mmask)
	}

	var sum int
	for _, v := range register {
		sum += v
	}

	fmt.Println(sum)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]Instruction, error) {
	file, err := os.Open("data/14.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rows []Instruction

	for scanner.Scan() {
		raw := scanner.Text()

		if strings.HasPrefix(raw, "mask") {
			rows = append(rows, Instruction{Mask: []byte(raw[6:])})
		} else {
			re := regexp.MustCompile(`^mem\[(\d*)]\s*=\s*(\d*)`)
			matches := re.FindAllStringSubmatch(raw, -1)

			addr, err := strconv.Atoi(matches[0][1])
			if err != nil {
				return nil, err
			}
			v, err := strconv.Atoi(matches[0][2])
			if err != nil {
				return nil, err
			}

			rows = append(rows, Instruction{Set: [2]int{addr, v}})
		}
	}

	return rows, nil
}
