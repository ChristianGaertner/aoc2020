package day04

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func (f Field) IsValid() bool {
	switch f.Type {
	case BYR:
		n, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}
		return n <= 2002 && n >= 1920
	case IYR:
		n, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}
		return n <= 2020 && n >= 2010
	case EYR:
		n, err := strconv.Atoi(f.Value)
		if err != nil {
			return false
		}
		return n <= 2030 && n >= 2020
	case HGT:
		if !strings.HasSuffix(f.Value, "cm") && !strings.HasSuffix(f.Value, "in") {
			return false
		}
		isCM := strings.HasSuffix(f.Value, "cm")
		n, err := strconv.Atoi(f.Value[:len(f.Value)-2])
		if err != nil {
			return false
		}

		if isCM {
			return n <= 193 && n >= 150
		}
		return n <= 76 && n >= 59
	case HCL:
		var re = regexp.MustCompile(`^#[0-9a-f]{6}$`)
		return re.MatchString(f.Value)
	case ECL:
		return f.Value == "amb" || f.Value == "blu" || f.Value == "brn" || f.Value == "gry" || f.Value == "grn" || f.Value == "hzl" || f.Value == "oth"
	case PID:
		var re = regexp.MustCompile(`^[0-9]{9}$`)
		return re.MatchString(f.Value)
	case CID:
		return true
	}
	panic("WHAT")
}

type Passport []Field

func (p Passport) HasAllFields() bool {
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

func (p Passport) IsValid() bool {
	if !p.HasAllFields() {
		return false
	}
	for _, f := range p {
		if !f.IsValid() {
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
		if p.HasAllFields() {
			numValid += 1
		}
	}

	fmt.Printf("numValid=%d\n", numValid)

	return nil
}

func SolvePartTwo() error {
	passports, err := _read()
	if err != nil {
		return err
	}

	var numValid int
	for _, p := range passports {
		if p.IsValid() {
			numValid += 1
		}
	}

	fmt.Printf("numValid=%d\n", numValid)

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
