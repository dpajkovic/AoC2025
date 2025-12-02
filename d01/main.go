package main

import (
	"bufio"
	"d01/p1"
	"d01/p2"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var input []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		input = append(input, l)
	}

	p1 := p1.P1(input)
	fmt.Printf("Result for part 1: %d\n", p1)
	p2 := p2.P2(input)
	fmt.Printf("Result for part 2: %d\n", p2)
}
