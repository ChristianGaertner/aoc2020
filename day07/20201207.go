package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Bag string

type Rule string

func (r Rule) CanHold(bag Bag) bool {
	return strings.Contains(string(r), string(bag)) && !strings.HasPrefix(string(r), string(bag))
}

func (r Rule) Holds() map[Bag]int {
	i := strings.Index(string(r), "bags contain ")
	relevant := strings.TrimSpace(string(r[i+len("bags contain ") : len(r)-1]))

	ret := make(map[Bag]int)

	for _, h := range strings.Split(relevant, ",") {
		h = strings.TrimSpace(h)
		if strings.Contains(h, "no other") {
			continue
		}
		num, err := strconv.Atoi(h[:1])
		if err != nil {
			fmt.Println(h)
			panic(err)
		}
		name := strings.TrimSpace(h[1 : len(h)-4])

		ret[Bag(name)] = num
	}

	return ret
}

func (r Rule) BagName() Bag {
	i := strings.Index(string(r), "bags contain")
	return Bag(r[:i-1])
}

type RuleTree struct {
	Bag     Bag
	Holders []RuleTree
}

func (t *RuleTree) Build(rules []Rule) {
	for _, r := range rules {
		if r.CanHold(t.Bag) {
			tree := RuleTree{
				Bag: r.BagName(),
			}
			tree.Build(rules)
			t.Holders = append(t.Holders, tree)
		}
	}
}

func (t *RuleTree) Bags() []Bag {
	bags := make(map[Bag]bool)
	bags[t.Bag] = true
	for _, c := range t.Holders {
		for _, b := range c.Bags() {
			bags[b] = true
		}
	}

	var out []Bag

	for b := range bags {
		out = append(out, b)
	}

	return out
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
	return "2020 12 07"
}

type ContainsMap map[Bag]map[Bag]int

func (contains ContainsMap) Count(bag Bag) int {
	var total int
	for child, num := range contains[bag] {
		total += contains.Count(child) * num
	}

	return total + 1
}

func SolvePartOne() error {
	rules, err := _read()
	if err != nil {
		return err
	}

	tree := RuleTree{
		Bag: "shiny gold",
	}
	tree.Build(rules)

	bags := tree.Bags()
	fmt.Println(len(bags) - 1)

	return nil
}

func SolvePartTwo() error {
	rules, err := _read()
	if err != nil {
		return err
	}

	contains := ContainsMap(make(map[Bag]map[Bag]int))

	for _, r := range rules {
		contains[r.BagName()] = r.Holds()
	}

	fmt.Println(contains.Count("shiny gold") - 1)

	return nil
}

func _read() ([]Rule, error) {
	file, err := os.Open("data/07.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var rules []Rule

	for scanner.Scan() {
		raw := scanner.Text()
		rules = append(rules, Rule(raw))
	}

	return rules, nil
}
