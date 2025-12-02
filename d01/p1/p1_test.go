package p1

import (
	"bufio"
	"os"
	"testing"
)

func TestP1(t *testing.T) {
	file, err := os.Open("../test_input.txt")
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
		want int
	}{
		{name: "Test with sample input",
			args: args{
				input: testInput,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := P1(tt.args.input); got != tt.want {
				t.Errorf("P1() = %v, want %v", got, tt.want)
			}
		})
	}
}
