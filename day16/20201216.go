package day16

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Range [2]int

type FieldRule struct {
	ValidRanges []Range
}

func (f FieldRule) Valid(t int) bool {
	for _, r := range f.ValidRanges {
		if t >= r[0] && t <= r[1] {
			return true
		}
	}

	return false
}

type Ticket []int

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
	rules, tickets, err := _read()
	if err != nil {
		return err
	}

	var s int

	for _, t := range tickets {
	t:
		for _, f := range t {
			for _, r := range rules {
				if r.Valid(f) {
					continue t
				}
			}
			s += f
		}
	}

	fmt.Println(s)

	return nil
}

func SolvePartTwo() error {
	_, _, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() (map[string]FieldRule, []Ticket, error) {
	file, err := os.Open("data/16.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[string]FieldRule)
	var tickets []Ticket

	m := 0

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}

		if raw == "your ticket:" {
			m++
			continue
		}
		if raw == "nearby tickets:" {
			m++
			continue
		}

		switch m {
		case 0:
			s, r, err := parseRule(raw)
			if err != nil {
				return nil, nil, err
			}
			rules[s] = r
		case 1:
			// ignore our ticket for now
			continue
		case 2:
			is, err := parseInts(raw)
			if err != nil {
				return nil, nil, err
			}
			tickets = append(tickets, is)
		}

	}

	return rules, tickets, nil
}

var re = regexp.MustCompile(`^(.*):\s(\d*)-(\d*)\sor\s(\d*)-(\d*)$`)

func parseRule(in string) (string, FieldRule, error) {
	var r FieldRule
	var name string
	for _, match := range re.FindAllStringSubmatch(in, -1) {
		name = match[1]

		ami, err := strconv.Atoi(match[2])
		ama, err := strconv.Atoi(match[3])

		bmi, err := strconv.Atoi(match[4])
		bma, err := strconv.Atoi(match[5])
		if err != nil {
			return "", FieldRule{}, err
		}

		r.ValidRanges = []Range{
			{ami, ama},
			{bmi, bma},
		}

	}
	return name, r, nil
}

func parseInts(in string) ([]int, error) {
	var res []int
	for _, i := range strings.Split(in, ",") {
		n, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		res = append(res, n)
	}
	return res, nil
}
