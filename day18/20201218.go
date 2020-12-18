package day18

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"strconv"
	"strings"
)

func EvalP1(e string) (int, int) {

	var l *int
	var r *int

	var op func(a, b int) int

	for pos := 0; pos < len(e); pos++ {
		t := e[pos]
		switch t {
		case ' ':
			continue
		case '+':
			op = func(a, b int) int {
				return a + b
			}
		case '-':
			op = func(a, b int) int {
				return a - b
			}
		case '*':
			op = func(a, b int) int {
				return a * b
			}
		case '/':
			op = func(a, b int) int {
				return a / b
			}
		case '(':
			num, length := EvalP1(e[pos+1:])
			pos += length
			if l == nil {
				l = &num
			} else {
				r = &num
			}
		case ')':
			if l == nil {
				panic("NIL ()")
			}
			return *l, pos + 1
		default:
			num, _ := strconv.Atoi(string(t))
			if l == nil {
				l = &num
			} else {
				r = &num
			}
		}

		if l != nil && r != nil {
			if op == nil {
				panic("NO OP")
			}
			num := op(*l, *r)
			l = &num
			r = nil
		}
	}

	if l == nil {
		panic("NIL RESULT")
	}

	return *l, 0
}

func EvalP2(e ast.Expr) int {

	switch e := e.(type) {
	case *ast.BinaryExpr:
		return EvalBinaryExpr(e)
	case *ast.BasicLit:
		switch e.Kind {
		case token.INT:
			i, _ := strconv.Atoi(e.Value)
			return i
		}
	case *ast.ParenExpr:
		return EvalP2(e.X)
	}

	return 0
}

func EvalBinaryExpr(exp *ast.BinaryExpr) int {
	left := EvalP2(exp.X)
	right := EvalP2(exp.Y)

	switch exp.Op {
	case token.ADD:
		return left * right
	case token.MUL:
		return left + right
	}

	return 0
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
	return "2020 12 18"
}

func SolvePartOne() error {
	expressions, err := _read()
	if err != nil {
		return err
	}

	var acc int
	for _, ex := range expressions {
		n, _ := EvalP1(ex)
		acc += n
	}

	fmt.Println(acc)

	return nil
}

func SolvePartTwo() error {
	expressions, err := _read2()
	if err != nil {
		return err
	}

	var acc int
	for _, ex := range expressions {
		n := EvalP2(ex)
		acc += n
	}

	fmt.Println(acc)

	return nil
}

func _read() ([]string, error) {
	file, err := os.Open("data/18.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expr []string

	for scanner.Scan() {
		raw := scanner.Text()

		if raw[0] == ';' {
			continue
		}

		expr = append(expr, raw)
	}

	return expr, nil
}

func _read2() ([]ast.Expr, error) {
	file, err := os.Open("data/18.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var expr []ast.Expr

	for scanner.Scan() {
		raw := scanner.Text()
		if raw[0] == ';' {
			continue
		}

		raw = strings.ReplaceAll(raw, "+", "$")
		raw = strings.ReplaceAll(raw, "*", "+")
		raw = strings.ReplaceAll(raw, "$", "*")

		tr, err := parser.ParseExpr(raw)
		if err != nil {
			return nil, err
		}
		expr = append(expr, tr)
	}

	return expr, nil
}
