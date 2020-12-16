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
	Name        string
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
	rules, tickets, _, err := _read()
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
	rules, tickets, myTicket, err := _read()
	if err != nil {
		return err
	}

	var validTickets []Ticket

o:
	for _, t := range tickets {
	t:
		for _, f := range t {
			for _, r := range rules {
				if r.Valid(f) {
					continue t
				}
			}
			continue o
		}
		validTickets = append(validTickets, t)
	}

	fi := make(map[string]int)
	locked := make(map[int]bool)

	for len(fi) != len(myTicket) {
		matches := make(map[string][]int)

		for _, r := range rules {
			for pos := 0; pos < len(myTicket); pos++ {
				if locked[pos] {
					continue
				}

				allValid := true
				for _, t := range validTickets {
					allValid = allValid && r.Valid(t[pos])
				}
				if allValid {
					matches[r.Name] = append(matches[r.Name], pos)
				}
			}

			if len(matches[r.Name]) == 1 {
				lockedIndex := matches[r.Name][0]
				fi[r.Name] = lockedIndex
				locked[lockedIndex] = true
			}

		}
	}

	acc := int64(1)

	for key, idx := range fi {
		if strings.HasPrefix(key, "departure") {
			acc *= int64(myTicket[idx])
		}
	}
	fmt.Println(acc)

	return nil
}

func _read() ([]FieldRule, []Ticket, Ticket, error) {
	file, err := os.Open("data/16.txt")
	if err != nil {
		return nil, nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules []FieldRule
	var tickets []Ticket
	var myTicket Ticket

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
			r, err := parseRule(raw)
			if err != nil {
				return nil, nil, nil, err
			}
			rules = append(rules, r)
		case 1:
			t, err := parseInts(raw)
			if err != nil {
				return nil, nil, nil, err
			}
			myTicket = t
		case 2:
			is, err := parseInts(raw)
			if err != nil {
				return nil, nil, nil, err
			}
			tickets = append(tickets, is)
		}

	}

	return rules, tickets, myTicket, nil
}

var re = regexp.MustCompile(`^(.*):\s(\d*)-(\d*)\sor\s(\d*)-(\d*)$`)

func parseRule(in string) (FieldRule, error) {
	var r FieldRule
	var name string
	for _, match := range re.FindAllStringSubmatch(in, -1) {
		name = match[1]

		ami, err := strconv.Atoi(match[2])
		ama, err := strconv.Atoi(match[3])

		bmi, err := strconv.Atoi(match[4])
		bma, err := strconv.Atoi(match[5])
		if err != nil {
			return r, err
		}

		r.ValidRanges = []Range{
			{ami, ama},
			{bmi, bma},
		}

	}
	r.Name = strings.Replace(name, " ", "_", 1)
	return r, nil
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

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
