package day19

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre"
	"os"
	"regexp"
	"strings"
)

type Rule struct {
	isLiteral  bool
	contentRaw string
	ruleNum    string
}

func (r Rule) GenerateRegex(rules map[string]Rule) string {
	if r.isLiteral {
		return r.contentRaw
	}

	res := "(?:"
	var ors []string
	for _, p := range strings.Split(r.contentRaw, " | ") {
		var c string
		for _, ruleNum := range strings.Split(p, " ") {
			c += rules[ruleNum].GenerateRegex(rules)
		}
		ors = append(ors, c)
	}

	return res + strings.Join(ors, "|") + ")"
}

func (r Rule) GenerateRegex2(rules map[string]Rule) string {
	if r.isLiteral {
		return r.contentRaw
	}

	res := "(?:"
	var ors []string
	for _, p := range strings.Split(r.contentRaw, " | ") {
		var c string
		for _, ruleNum := range strings.Split(p, " ") {
			if ruleNum == "8" {
				del := rules["42"].GenerateRegex2(rules)
				c += del + "+"
			} else if ruleNum == "11" {
				del := rules["42"].GenerateRegex2(rules)
				notDel := rules["31"].GenerateRegex2(rules)
				c += fmt.Sprintf("(?<m>(%s%s|%[1]s(?&m)%[2]s))", del, notDel)
			} else {
				c += rules[ruleNum].GenerateRegex(rules)
			}
		}
		ors = append(ors, c)
	}

	return res + strings.Join(ors, "|") + ")"
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
	return "2020 12 19"
}

func SolvePartOne() error {
	rules, messages, err := _read()
	if err != nil {
		return err
	}

	res := rules["0"].GenerateRegex(rules)
	re := regexp.MustCompile("^" + res + "$")

	var i int

	for _, msg := range messages {
		if re.MatchString(msg) {
			i++
		}
	}

	fmt.Println(i)

	return nil
}

func SolvePartTwo() error {
	rules, messages, err := _read()
	if err != nil {
		return err
	}

	rules["8"] = parseRule("8: 42 | 42 8")
	rules["11"] = parseRule("11: 42 31 | 42 11 31")

	res := rules["0"].GenerateRegex2(rules)
	re, cerr := pcre.Compile("^"+res+"$", 0)
	if cerr != nil {
		return errors.New(cerr.Message)
	}

	var i int

	for _, msg := range messages {
		if re.MatcherString(msg, 0).Matches() {
			i++
		}
	}

	fmt.Println(i)

	return nil
}

func _read() (map[string]Rule, []string, error) {
	file, err := os.Open("data/19.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	rules := make(map[string]Rule)
	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			break
		}
		if raw[0] == ';' {
			continue
		}
		r := parseRule(raw)
		rules[r.ruleNum] = r
	}

	var messages []string
	for scanner.Scan() {
		raw := scanner.Text()
		if raw[0] == ';' {
			continue
		}

		messages = append(messages, raw)
	}

	return rules, messages, nil
}

func parseRule(in string) Rule {
	strSplit := strings.SplitN(in, ": ", 2)
	if strSplit[1][0] == '"' {
		return Rule{
			ruleNum:    strSplit[0],
			contentRaw: strSplit[1][1 : len(strSplit[1])-1],
			isLiteral:  true,
		}
	}
	return Rule{
		ruleNum:    strSplit[0],
		contentRaw: strSplit[1],
	}
}
