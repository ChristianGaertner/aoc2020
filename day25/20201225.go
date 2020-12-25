package day25

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
	return "2020 12 25"
}

func SolvePartOne() error {
	card, door, err := _read()
	if err != nil {
		return err
	}

	var loop int
	for i := 1; i != card; loop++ {
		i = i * 7 % 20201227
	}

	key := 1
	for l := 0; l < loop; l++ {
		key = key * door % 20201227
	}

	fmt.Println(key)

	return nil
}

func SolvePartTwo() error {
	_, _, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() (int, int, error) {
	file, err := os.Open("data/25.txt")
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	card, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, 0, err
	}
	scanner.Scan()
	door, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return 0, 0, err
	}
	return card, door, nil
}
