package day13

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
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
	return "2020 12 13"
}

func SolvePartOne() error {
	ts, buses, err := _read()
	if err != nil {
		return err
	}

	var busID int
	diff := math.MaxInt32

	for _, b := range buses {
		if b == -1 {
			continue
		}
		for i := 0; i < ts*2; i++ {
			t := b * i
			if t < ts {
				continue
			} else {
				d := t - ts
				if d < diff {
					diff = d
					busID = b
				}
				break
			}
		}
	}

	fmt.Println(busID * diff)

	return nil
}

func SolvePartTwo() error {
	_, buses, err := _read()
	if err != nil {
		return err
	}

	pos := 0
	offset := 1

	for i, b := range buses {
		if b == -1 {
			continue
		}
		for ((pos + i) % b) != 0 {
			pos += offset
		}
		offset *= b
	}

	fmt.Println(pos)

	return nil
}

func _read() (int, []int, error) {
	file, err := os.Open("data/13.txt")
	if err != nil {
		return 0, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	ts, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, nil, err
	}

	var buses []int

	scanner.Scan()
	for _, bus := range strings.Split(scanner.Text(), ",") {
		if bus == "x" {
			buses = append(buses, -1)
			continue
		}
		b, err := strconv.Atoi(bus)
		if err != nil {
			return 0, nil, err
		}
		buses = append(buses, b)
	}

	return ts, buses, nil
}
