//	Copyright (c) Miloš Rackov 2025
//	Distributed under the Boost Software License, Version 1.0.
//	(See accompanying file LICENSE or copy at
//	https://www.boost.org/LICENSE_1_0.txt)

package p2

import (
	"bufio"
	"os"
	"testing"
)

const testResult = "3263827"

func TestP2(t *testing.T) {
	file, err := os.Open("../testinput.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var testInput []string

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		l := scanner.Text()
		testInput = append(testInput, l)
	}
	type args struct {
		input []string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "Test with sample input",
			args: args{
				input: testInput,
			},
			want: testResult,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P2(tt.args.input); got != tt.want {
				t.Errorf("P2() = %v, want %v", got, tt.want)
			}
		})
	}
}
