package day09

import (
	"bufio"
	"errors"
	"fmt"
	"math"
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

func FindInvalid(windowLen int, numbers []int64) int64 {
	for i, n := range numbers {
		if i < windowLen {
			continue
		}
		window := numbers[i-windowLen : i]
		if !IsValid(n, window) {
			return n
		}
	}
	return -1
}

func SolvePartOne() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	fmt.Println(FindInvalid(25, numbers))

	return nil
}

func SolvePartTwo() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	invalid := FindInvalid(25, numbers)
	if invalid == -1 {
		return errors.New("no invalid number found")
	}

	var tally int64
	tallyStart := 0

	var contSet []int64

	for i := 0; i < len(numbers); {
		tally += numbers[i]

		if tally == invalid {
			contSet = numbers[tallyStart : i+1]
			break
		}

		if tally < invalid {
			// continue
			i++
		} else {
			// go back
			i = tallyStart + 1
			tally = 0
			tallyStart++
		}
	}

	if contSet == nil {
		return errors.New("no contiguous set found")
	}

	var sum int64
	min := int64(math.MaxInt64)
	max := int64(math.MinInt64)
	for _, n := range contSet {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
		sum += n
	}

	if sum != invalid {
		return errors.New("check failed sum != invalid")
	}

	fmt.Printf("min=%d\tmax=%d\tsolution=%d\n", min, max, min+max)

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
