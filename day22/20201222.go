package day22

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Stack []int

func (s *Stack) Pop() int {
	v := (*s)[0]
	*s = (*s)[1:]
	return v
}

func (s *Stack) Add(v ...int) {
	*s = append(*s, v...)
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
	return "2020 12 22"
}

func SolvePartOne() error {
	cardsA, cardsB, err := _read()
	if err != nil {
		return err
	}

	var i int
	for len(cardsA) != 0 && len(cardsB) != 0 {
		i++

		//fmt.Printf("-- Round %d --\n", i)
		//fmt.Printf("Player 1's deck: %v\n", cardsA)
		//fmt.Printf("Player 2's deck: %v\n", cardsB)

		a, b := cardsA.Pop(), cardsB.Pop()

		//fmt.Printf("Player 1's plays: %v\n", a)
		//fmt.Printf("Player 2's plays: %v\n", b)

		if a > b {
			//fmt.Printf("Player 1 wins the round!\n")
			cardsA.Add(a, b)
		} else {
			//fmt.Printf("Player 2 wins the round!\n")
			cardsB.Add(b, a)
		}
	}

	s := cardsA
	if len(s) == 0 {
		s = cardsB
	}

	var score int

	for i, v := range s {
		score += v * (len(s) - i)
	}

	fmt.Println(score)

	return nil
}

func SolvePartTwo() error {
	_, _, err := _read()
	if err != nil {
		return err
	}

	return nil
}

func _read() (Stack, Stack, error) {
	file, err := os.Open("data/22.txt")
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var cardA []int
	var cardB []int

	var mode int

	for scanner.Scan() {
		raw := scanner.Text()
		if raw == "" {
			continue
		}
		if raw[0] == ';' {
			continue
		}
		if raw[0] == 'P' {
			mode++
			continue
		}

		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, nil, err
		}

		if mode == 1 {
			cardA = append(cardA, n)
		} else {
			cardB = append(cardB, n)
		}
	}

	return cardA, cardB, nil
}
