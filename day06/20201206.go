package day06

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var letters = [...]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

type Group []string

func (g Group) Yes() []string {
	var yes []string

outer:
	for _, l := range letters {
		for _, answer := range g {
			if strings.ContainsRune(answer, l) {
				yes = append(yes, strconv.QuoteRune(l))
				continue outer
			}
		}
	}
	return yes
}

func (g Group) UniqueYes() []string {
	var uniqueYes []string

	for _, l := range letters {
		allYes := true
		for _, answer := range g {
			allYes = allYes && strings.ContainsRune(answer, l)
		}
		if allYes {
			uniqueYes = append(uniqueYes, strconv.QuoteRune(l))
		}
	}
	return uniqueYes
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
	return "2020 12 06"
}

func SolvePartOne() error {
	groups, err := _read()
	if err != nil {
		return err
	}

	var total int
	for _, g := range groups {
		n := len(g.Yes())
		total += n
	}

	fmt.Printf("total=%d\n", total)

	return nil
}

func SolvePartTwo() error {
	groups, err := _read()
	if err != nil {
		return err
	}

	var total int
	for _, g := range groups {
		n := len(g.UniqueYes())
		total += n
	}

	fmt.Printf("total=%d\n", total)

	return nil
}

func _read() ([]Group, error) {
	file, err := os.Open("data/06.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var groups []Group

	var current Group

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			groups = append(groups, current)
			current = nil
			continue
		}
		current = append(current, raw)
	}

	groups = append(groups, current)
	return groups, nil
}
