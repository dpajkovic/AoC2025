//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"bufio"
	"d03/p1"
	"d03/p2"
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
	fmt.Printf("Result for part 1: %s\n", p1)
	p2 := p2.P2(input)
	fmt.Printf("Result for part 2: %s\n", p2)
}
