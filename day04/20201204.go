package day04

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type FieldType string

const (
	BYR FieldType = "byr"
	IYR FieldType = "iyr"
	EYR FieldType = "eyr"
	HGT FieldType = "hgt"
	HCL FieldType = "hcl"
	ECL FieldType = "ecl"
	PID FieldType = "pid"
	CID FieldType = "cid"
)

var requiredFields = []FieldType{BYR, IYR, EYR, HGT, HCL, ECL, PID}

type Field struct {
	Type  FieldType
	Value string
}

type Passport []Field

func (p Passport) IsValid() bool {
	for _, req := range requiredFields {
		found := false
		for _, f := range p {
			if f.Type == req {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
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
	return "2020 12 04"
}

func SolvePartOne() error {
	passports, err := _read()
	if err != nil {
		return err
	}

	var numValid int
	for _, p := range passports {
		fmt.Println(p)
		if p.IsValid() {
			numValid += 1
		}
	}

	fmt.Printf("numValid=%d\n", numValid)

	return nil
}

func SolvePartTwo() error {
	_, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() ([]Passport, error) {
	file, err := os.Open("data/04.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var passports []Passport

	scanner := bufio.NewScanner(file)

	var current Passport
	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			passports = append(passports, current)
			current = Passport{}
			continue
		}
		fieldPairs := strings.Split(raw, " ")
		for _, p := range fieldPairs {
			kv := strings.SplitN(p, ":", 2)
			field := Field{
				Type:  FieldType(kv[0]),
				Value: kv[1],
			}
			current = append(current, field)
		}
	}

	passports = append(passports, current)

	return passports, nil
}
