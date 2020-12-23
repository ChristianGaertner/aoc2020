package day23

import (
	"bufio"
	"container/ring"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Solver struct{}

func (Solver) Solve() error {
	err := SolvePartOne()
	if err != nil {
		return err
	}
	return SolvePartTwo()
}

func (Solver) Day() string {
	return "2020 12 23"
}

func contains(head *ring.Ring, value int) bool {
	current := head
	for i := 0; i < current.Len(); i++ {
		if current.Value.(int) == value {
			return true
		}
		current = current.Next()
	}
	return false
}

func max(r *ring.Ring) int {
	max := 0
	r.Do(func(i interface{}) {
		num, ok := i.(int)
		if ok {
			if num > max {
				max = num
			}
		}
	})
	return max
}

func appen(head *ring.Ring, value int) *ring.Ring {
	current := head
	for i := 0; i < head.Len(); i++ {
		if current.Value.(int) == value {
			return current
		}
		current = current.Next()
	}
	return nil
}

func play(in []int) string {
	cups := ring.New(len(in))
	for _, i := range in {
		cups.Value = i
		cups = cups.Next()
	}

	for i := 0; i < 100; i++ {
		picked := cups.Unlink(3)
		dst := cups.Value.(int) - 1
		if dst == 0 {
			dst = max(cups)
		}

		for contains(picked, dst) {
			dst--
			if dst == 0 {
				dst = max(cups)
			}
		}
		dstRing := appen(cups, dst)
		dstRing.Link(picked)
		cups = cups.Next()
	}

	var answer string

	appen(cups, 1).Next().Do(func(i interface{}) {
		if i.(int) == 1 {
			return
		}
		answer += fmt.Sprint(i)
	})
	return answer
}

func play2(size int, in []int) int {
	cups := ring.New(size)
	overflow := make([]*ring.Ring, size)
	for _, i := range in {
		cups.Value = i
		overflow[i-1] = cups
		cups = cups.Next()
	}

	maxN := 0
	for i := max(cups) + 1; i <= size; i++ {
		cups.Value = i
		overflow[i-1] = cups
		maxN = i
		cups = cups.Next()
	}

	for i := 0; i < size; i++ {
		picked := cups.Unlink(3)
		dst := cups.Value.(int) - 1
		if dst == 0 {
			dst = max(cups)
		}

		for contains(picked, dst) {
			dst--
			if dst == 0 {
				dst = maxN
			}

		}
		dstRing := overflow[dst-1]
		dstRing.Link(picked)
		cups = cups.Next()
	}

	a := overflow[0]
	b := a.Next()
	c := b.Next()

	return b.Value.(int) * c.Value.(int)
}

func SolvePartOne() error {
	labels, err := _read()
	if err != nil {
		return err
	}

	fmt.Println(play(labels))

	return nil
}

func SolvePartTwo() error {
	labels, err := _read()
	if err != nil {
		return err
	}

	fmt.Println(play2(1000000, labels))

	return nil
}

func _read() ([]int, error) {
	file, err := os.Open("data/23.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}
		if raw[0] == ';' {
			continue
		}

		var res []int
		for _, r := range strings.Split(raw, "") {
			n, err := strconv.Atoi(r)
			if err != nil {
				return nil, err
			}
			res = append(res, n)
		}
		return res, nil
	}
	return nil, errors.New("no input")
}
