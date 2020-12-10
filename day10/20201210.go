package day10

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Solver struct{}

type Adapter struct {
	Rating   int
	IsDevice bool
}

func DeviceFromAdapter(as []Adapter) Adapter {
	max := math.MinInt32

	for _, a := range as {
		if a.Rating > max {
			max = a.Rating
		}
	}

	return Adapter{
		Rating:   max + 3,
		IsDevice: true,
	}
}

func FindNextAdapter(jolts int, as []Adapter) (Adapter, []Adapter) {
	var minIndex int
	minAdapter := Adapter{
		Rating: math.MaxInt32,
	}

	for i, a := range as {
		if jolts+3 < a.Rating {
			continue
		}
		if a.Rating < minAdapter.Rating {
			minIndex = i
			minAdapter = a
		}
	}

	next := append(as[:minIndex], as[minIndex+1:]...)

	return minAdapter, next
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

	currentJolts := 0

	toChain := adapter[:]

	var a Adapter
	var diffs []int
	for range adapter {
		a, toChain = FindNextAdapter(currentJolts, toChain)
		diffs = append(diffs, a.Rating-currentJolts)
		currentJolts = a.Rating
	}

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
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]Adapter, error) {
	file, err := os.Open("data/10.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []Adapter

	for scanner.Scan() {
		raw := scanner.Text()
		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		lines = append(lines, Adapter{Rating: n})
	}

	return lines, nil
}
