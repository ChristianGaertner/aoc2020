package day02

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Solver struct{}

type Policy struct {
	Min    int
	Max    int
	Letter string
}

type PolicyAndPassword struct {
	Policy   Policy
	Password string
}

func isValidPartOne(p PolicyAndPassword) bool {
	tester := regexp.MustCompile(p.Policy.Letter)

	matches := tester.FindAllStringIndex(p.Password, -1)

	numMatches := len(matches)

	if numMatches < p.Policy.Min {
		return false
	}
	if numMatches > p.Policy.Max {
		return false
	}
	return true
}

func isValidPartTwo(p PolicyAndPassword) bool {

	var posMatches int

	for i, r := range p.Password {
		pos := i + 1
		letter := string([]rune{r})
		if (pos == p.Policy.Min || pos == p.Policy.Max) && letter == p.Policy.Letter {
			posMatches += 1
		}
	}

	return posMatches == 1
}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 02"
}

func SolvePartOne() error {
	passwords, err := _read()
	if err != nil {
		return err
	}

	var numValid int

	fmt.Println("Part One")
	for _, pass := range passwords {
		if isValidPartOne(pass) {
			numValid += 1
		}
	}

	fmt.Printf("num Valid=%d\n", numValid)
	return nil
}

func SolvePartTwo() error {
	passwords, err := _read()
	if err != nil {
		return err
	}

	var numValid int

	fmt.Println("Part Two")
	for _, pass := range passwords {
		if isValidPartTwo(pass) {
			numValid += 1
		}
	}

	fmt.Printf("num Valid=%d\n", numValid)
	return nil
}

func _read() ([]PolicyAndPassword, error) {
	file, err := os.Open("data/02.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pap []PolicyAndPassword

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		n, err := _parseLine(raw)
		if err != nil {
			return nil, err
		}
		pap = append(pap, n)
	}

	return pap, nil
}

var re = regexp.MustCompile(`^(\d*)-(\d*)\s([a-zA-z]):\s(.*)$`)

func _parseLine(l string) (PolicyAndPassword, error) {
	pap := PolicyAndPassword{}
	for _, match := range re.FindAllStringSubmatch(l, -1) {
		for i, x := range match {
			var err error
			if i == 0 {
				continue
			} else if i == 1 {
				pap.Policy.Min, err = strconv.Atoi(x)
				if err != nil {
					return PolicyAndPassword{}, err
				}
			} else if i == 2 {
				pap.Policy.Max, err = strconv.Atoi(x)
				if err != nil {
					return PolicyAndPassword{}, err
				}
			} else if i == 3 {
				pap.Policy.Letter = x
			} else if i == 4 {
				pap.Password = x
			}
		}
	}
	return pap, nil
}
