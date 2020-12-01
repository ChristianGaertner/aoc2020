package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

const _sumTarget = 2020

func Solve20201201() error {
	numbers, err := _read()
	if err != nil {
		return err
	}

	for i, a := range numbers {
		for _, b := range numbers[i:] {
			if a + b == _sumTarget {
				fmt.Println(a * b)
			}
		}
	}
	return nil
}

func _read() ([]int, error) {
	file, err := os.Open("data/01.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		raw := scanner.Text()

		n, err := strconv.Atoi(raw)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, n)
	}

	sort.Ints(numbers)
	return numbers, nil
}