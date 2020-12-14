package day14

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Mask [2]int

type Instruction struct {
	Mask   Mask
	IsMask bool
	Set    [2]int
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
		if ins.IsMask {
			mask = ins.Mask[0]
			mmask = ins.Mask[1]
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

func SetAll(register map[int]int, addr int, rm int, value int) {
	if rm == 0 {
		register[addr] = value
		return
	}
	b := rm & -rm
	rm = rm &^ b
	SetAll(register, addr, rm, value)
	SetAll(register, addr|b, rm, value)
}

func SolvePartTwo() error {
	instructions, err := _read()
	if err != nil {
		return err
	}

	register := make(map[int]int)

	mask := 0
	mmask := 0

	for _, ins := range instructions {
		if ins.IsMask {
			mask = ins.Mask[0]
			mmask = ins.Mask[1]
			continue
		}

		addr := ins.Set[0]
		v := ins.Set[1]

		addr = (addr & mmask) | mask
		SetAll(register, addr, ((1<<36)-1)&^mmask, v)
	}

	var sum int
	for _, v := range register {
		sum += v
	}

	fmt.Println(sum)

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
			mask := 0
			mmask := 0
			for _, c := range []byte(raw[6:]) {
				mask = mask << 1
				mmask = mmask << 1
				if c != 'X' {
					mmask |= 1
					if c == '1' {
						mask |= 1
					}
				}
			}

			rows = append(rows, Instruction{IsMask: true, Mask: [2]int{
				mask, mmask,
			}})
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
