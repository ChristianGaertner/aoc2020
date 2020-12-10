package day10

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Solver struct{}

type Rating int

func DeviceFromAdapter(as []Rating) Rating {
	max := math.MinInt32

	for _, a := range as {
		if int(a) > max {
			max = int(a)
		}
	}

	return Rating(max + 3)
}

func Diffs(as []Rating) []int {
	sort.Slice(as, func(i, j int) bool {
		return as[i] < as[j]
	})
	var diffs []int
	var c Rating
	for _, a := range as {
		diffs = append(diffs, int(a-c))
		c = a
	}
	return diffs
}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 10"
}

func SolvePartOne() error {
	adapter, err := _read()
	if err != nil {
		return err
	}

	adapter = append(adapter, DeviceFromAdapter(adapter))
	diffs := Diffs(adapter)

	var numOneDiff int
	var numThreeDiff int
	for _, d := range diffs {
		if d == 1 {
			numOneDiff++
		} else if d == 3 {
			numThreeDiff++
		} else {
			return errors.New("unexpected diff")
		}
	}

	fmt.Println(numOneDiff * numThreeDiff)

	return nil
}

func SolvePartTwo() error {
	adapter, err := _read()
	if err != nil {
		return err
	}

	adapter = append(adapter, DeviceFromAdapter(adapter))
	sort.Slice(adapter, func(i, j int) bool {
		return adapter[i] < adapter[j]
	})

	acc := map[Rating]int{0: 1}

	for _, a := range adapter {
		acc[a] = acc[a-1] + acc[a-2] + acc[a-3]
	}

	fmt.Println(acc[adapter[len(adapter)-1]])

	return nil
}

func _read() ([]Rating, error) {
	file, err := os.Open("data/10.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []Rating

	for scanner.Scan() {
		raw := scanner.Text()
		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		lines = append(lines, Rating(n))
	}

	return lines, nil
}
