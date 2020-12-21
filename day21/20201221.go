package day21

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Meal struct {
	Ing       []string
	Allergies []string
}

type Set map[string]bool

func NewSet() Set {
	return make(Set)
}

func NewSetFrom(o Set) Set {
	r := NewSet()
	for v := range o {
		r.Add(v)
	}
	return r
}

func (s Set) Single() string {
	if len(s) != 1 {
		panic("NOP")
	}
	for v := range s {
		return v
	}
	panic("NOP")
}

func (s Set) Add(ts string) {
	s[ts] = true
}

func (s Set) AddAll(ts []string) {
	for _, t := range ts {
		s[t] = true
	}
}

func (s Set) Contains(ts string) bool {
	_, ok := s[ts]
	return ok
}

func (s Set) Remove(ts string) {
	delete(s, ts)
}

func (s Set) RemoveAll(ts []string) {
	for _, t := range ts {
		delete(s, t)
	}
}

func (s Set) Subtract(o Set) Set {
	r := NewSetFrom(s)
	for v := range o {
		r.Remove(v)
	}
	return r
}

func (s Set) Union(o Set) Set {
	r := NewSet()
	for ts := range s {
		r.Add(ts)
	}
	for ts := range o {
		r.Add(ts)
	}
	return r
}

func (s Set) Intersect(o Set) Set {
	r := NewSet()
	for ts := range s {
		if o.Contains(ts) {
			r.Add(ts)
		}
	}
	return r
}

func FindAll(meals []Meal) (Set, Set) {
	allIngredients := NewSet()
	allAllergens := NewSet()

	for _, m := range meals {
		allIngredients.AddAll(m.Ing)
		allAllergens.AddAll(m.Allergies)
	}
	return allIngredients, allAllergens
}

func FindCommon(meals []Meal) map[string]Set {
	_, allAllergens := FindAll(meals)
	common := make(map[string]Set)

	for a := range allAllergens {
		s := NewSet()
		for _, meal := range meals {
			if contains(meal.Allergies, a) {
				ss := NewSet()
				ss.AddAll(meal.Ing)
				if len(s) == 0 {
					s = s.Union(ss)
				} else {
					s = s.Intersect(ss)
				}
			}
		}
		common[a] = s
	}
	return common
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
	return "2020 12 21"
}

func SolvePartOne() error {
	meals, err := _read()
	if err != nil {
		return err
	}

	allIngredients, _ := FindAll(meals)
	common := FindCommon(meals)

	haveIng := NewSet()
	for _, c := range common {
		haveIng = haveIng.Union(c)
	}

	noHave := allIngredients.Subtract(haveIng)

	var acc int

	for _, m := range meals {
		for _, i := range m.Ing {
			if noHave.Contains(i) {
				acc++
			}
		}
	}

	fmt.Println(acc)

	return nil
}

func SolvePartTwo() error {
	meals, err := _read()
	if err != nil {
		return err
	}

	allIngredients, allAllergens := FindAll(meals)
	common := FindCommon(meals)

	haveIng := NewSet()
	for _, c := range common {
		haveIng = haveIng.Union(c)
	}

	noHave := allIngredients.Subtract(haveIng)

	allergenMap := make(map[string]Set)

	for a := range allAllergens {
		s := NewSet()
		for _, m := range meals {
			if contains(m.Allergies, a) {
				ss := NewSet()
				ss.AddAll(m.Ing)
				ss = ss.Subtract(noHave)
				if len(s) == 0 {
					s = s.Union(ss)
				} else {
					s = s.Intersect(ss)
				}
			}
		}
		if len(s) > 0 {
			allergenMap[a] = s
		}
	}

	res := make(map[string]string)

	for len(allergenMap) > 0 {
		toRemove := NewSet()
		for k, v := range allergenMap {
			if len(v) == 1 {
				res[k] = v.Single()
				toRemove.Add(v.Single())
			}
		}
		for k, v := range allergenMap {
			allergenMap[k] = v.Subtract(toRemove)
			if len(allergenMap[k]) == 0 {
				delete(allergenMap, k)
			}
		}
	}

	var fs []string
	for k := range res {
		fs = append(fs, k)
	}
	sort.Strings(fs)

	var answer string
	for _, a := range fs {
		answer += res[a] + ","
	}

	fmt.Println(strings.TrimRight(answer, ","))

	return nil
}

func _read() ([]Meal, error) {
	file, err := os.Open("data/21.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var meals []Meal

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}
		if raw[0] == ';' {
			continue
		}
		in := strings.Split(strings.Split(raw, " (")[0], " ")

		alls := strings.Split(strings.Split(raw, "(contains ")[1], ", ")

		alls[len(alls)-1] = strings.TrimRight(alls[len(alls)-1], ")")

		meals = append(meals, Meal{
			Ing:       in,
			Allergies: alls,
		})
	}

	return meals, nil
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
