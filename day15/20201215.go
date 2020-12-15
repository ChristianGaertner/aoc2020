package day15

import (
	"bufio"
	"fmt"
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
	return "2020 12 1"
}

func SolvePartOne() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	last := make(map[int]int)
	diffs := make(map[int]int)

	var lastNumber int
	for t := 1; t <= 2020; t++ {
		if t <= len(numbers) {
			lastNumber = numbers[t-1]
			v, ok := last[numbers[t-1]]
			last[lastNumber] = t
			if ok {
				diffs[lastNumber] = t - v
			}
			continue
		}

		v, ok := diffs[lastNumber]
		if ok {
			lastNumber = v
		} else {
			lastNumber = 0
		}
		v, ok = last[lastNumber]
		if ok {
			diffs[lastNumber] = t - v
		}
		last[lastNumber] = t
	}

	fmt.Println(lastNumber)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]int, error) {
	file, err := os.Open("data/15.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	raw := scanner.Text()

	var ints []int
	for _, r := range strings.Split(raw, ",") {
		n, err := strconv.Atoi(r)
		if err != nil {
			return nil, err
		}
		ints = append(ints, n)
	}

	return ints, nil
}
