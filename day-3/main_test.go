package main

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func Test_getValidNumbers(t *testing.T) {
	tests := map[string]struct {
		in   [][]int32
		want []int
	}{
		"467 valid; 114 invalid": {
			in: [][]int32{
				{4, 6, 7, dotVal, dotVal, 1, 1, 4, dotVal, dotVal},
				{dotVal, dotVal, dotVal, specialVal, dotVal, dotVal, dotVal, dotVal, dotVal, dotVal},
			},
			want: []int{467},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := getValidNumbers(tt.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getValidNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSumOfValidNumers(t *testing.T) {
	tests := map[string]struct {
		input io.ReadCloser
		want  int
	}{
		"sum with valid number at end of line": {
			input: fromFile("testfiles/sum_10428.txt"),
			want:  10428,
		},
		"example sum": {
			input: fromFile("testfiles/sum_4361.txt"),
			want:  4361,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			defer tt.input.Close()
			inputArray := inputTo2DArray(tt.input)

			if got := getSumOfValidNumers(inputArray); got != tt.want {
				t.Errorf("getSumOfValidNumers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func fromFile(name string) *os.File {
	file, err := os.Open(name)
	if err != nil {
		panic(err)
	}

	return file
}
