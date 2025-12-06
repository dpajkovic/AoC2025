//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"d05/p1"
	"d05/p2"
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

	t := time.Now()
	p1 := p1.P1(input)
	fmt.Printf("Result for part 1: %s\nExecution time: %s\n", p1, time.Since(t))
	t = time.Now()
	p2 := p2.P2(input)
	fmt.Printf("Result for part 1: %s\nExecution time: %s\n", p2, time.Since(t))
}
