package dayone

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const _sumTarget = 2020

type Solver struct {}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 01";
}

func SolvePartOne() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	fmt.Println("Part One")
	for i, a := range numbers {
		for _, b := range numbers[i:] {
			if a + b == _sumTarget {
				fmt.Println(a * b)
			}
		}
	}
	return nil
}

func SolvePartTwo() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	fmt.Println("Part Two")
	for i, a := range numbers {
		for j, b := range numbers[i:] {
			for _, c := range numbers[j:] {
				if a + b + c == _sumTarget {
					fmt.Println(a * b * c)
				}
			}
		}
	}
	return nil
}

func _read() ([]int, error) {
	file, err := os.Open("data/01.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	return numbers, nil
}