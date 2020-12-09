package day09

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	return "2020 12 09"
}

func IsValid(n int64, window []int64) bool {

	for i, a := range window {
		for _, b := range window[i:] {
			if a+b == n {
				return true
			}
		}
	}
	return false
}

func SolvePartOne() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	windowLen := 25

	for i, n := range numbers {
		if i < windowLen {
			continue
		}
		window := numbers[i-windowLen : i]
		if !IsValid(n, window) {
			fmt.Println(n)
		}
	}

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]int64, error) {
	file, err := os.Open("data/09.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []int64

	for scanner.Scan() {
		raw := scanner.Text()
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return nil, err
		}
		lines = append(lines, n)
	}

	return lines, nil
}
