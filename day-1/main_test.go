package main

import (
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_replace(t *testing.T) {
	type args struct {
		line string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"tdszrfzspthree2ttzseven5seven": {
			args: args{
				line: "tdszrfzspthree2ttzseven5seven",
			},
			want: "3257",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := replace(tt.args.line); got != tt.want {
				t.Errorf("repalceWrittenNumbersWithLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repalceWrittenNumbersWithLetters(t *testing.T) {
	type args struct {
		line string
	}
	tests := map[string]struct {
		args args
		want string
	}{
		"tdszrfzspthree2ttzseven5seven": {
			args: args{
				line: "tdszrfzspthree2ttzseven5seven",
			},
			want: "tdszrfzsp32ttzseven57",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := replaceWrittenNumbersWithLetters(tt.args.line)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_sumFromIoReader(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := map[string]struct {
		args args
		want int
	}{
		"sample": {
			args: args{
				reader: strings.NewReader(`two1nine
				eightwothree
				abcone2threexyz
				xtwone3four
				4nineeightseven2
				zoneight234
				7pqrstsixteen`),
			},
			want: 281,
		},
		"other": {
			args: args{
				reader: strings.NewReader(`mkqtlrzmzfsix2ccqsnnxtwo4sevenxp9
				tdszrfzspthree2ttzseven5seven
				four14three7
				4fdkcclmxmxsevenfiver
				1`),
			},
			want: 69 + 37 + 47 + 45 + 11,
		},
		"the reason why i neeeded to cheat": {
			args: args{
				reader: strings.NewReader(`88eightwoffg`),
			},
			want: 82,
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := sumFromIoReader(tt.args.reader)
			assert.Nil(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
