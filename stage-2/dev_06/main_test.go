package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_cut(t *testing.T) {
	for _, testCase := range []struct {
		name    string
		flags   flagsStrusct
		inFile  string
		outFile string
	}{
		{
			name:   `clear && cat input-01.txt | cut -f1 -d: -s && echo "===" && cat input-01.txt | go run main.go -f=1 -d=: -s=true`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f1-s.txt`,
			flags: flagsStrusct{f: 1, s: true, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f2 -d: -s && echo "===" && cat input-01.txt | go run main.go -f=2 -d=: -s=true`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f2-s.txt`,
			flags: flagsStrusct{f: 2, s: true, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f3 -d: -s && echo "===" && cat input-01.txt | go run main.go -f=3 -d=: -s=true`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f3-s.txt`,
			flags: flagsStrusct{f: 3, s: true, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f4 -d: -s && echo "===" && cat input-01.txt | go run main.go -f=4 -d=: -s=true`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f4-s.txt`,
			flags: flagsStrusct{f: 4, s: true, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f5 -d: -s && echo "===" && cat input-01.txt | go run main.go -f=5 -d=: -s=true`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f5-s.txt`,
			flags: flagsStrusct{f: 5, s: true, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f1 -d: && echo "===" && cat input-01.txt | go run main.go -f=1 -d=:`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f1.txt`,
			flags: flagsStrusct{f: 1, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f2 -d: && echo "===" && cat input-01.txt | go run main.go -f=2 -d=:`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f2.txt`,
			flags: flagsStrusct{f: 2, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f3 -d: && echo "===" && cat input-01.txt | go run main.go -f=3 -d=:`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f3.txt`,
			flags: flagsStrusct{f: 3, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f4 -d: && echo "===" && cat input-01.txt | go run main.go -f=4 -d=:`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f4.txt`,
			flags: flagsStrusct{f: 4, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f5 -d: && echo "===" && cat input-01.txt | go run main.go -f=5 -d=:`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d-f5.txt`,
			flags: flagsStrusct{f: 5, d: ':'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f1 -d1 && echo "===" && cat input-01.txt | go run main.go -f=1 -d=1`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d1-f1.txt`,
			flags: flagsStrusct{f: 1, d: '1'},
		},
		{
			name:   `clear && cat input-01.txt | cut -f2 -d1  && echo "===" && cat input-01.txt | go run main.go -f=2 -d=1`,
			inFile: `input-01.txt`, outFile: `input-01-cut-d1-f2.txt`,
			flags: flagsStrusct{f: 2, d: '1'},
		},
	} {

		input, _ := getRowsFromFile(testCase.inFile)
		output, _ := getRowsFromFile(testCase.outFile)

		s := cutStruct{
			flags: testCase.flags,
		}

		result := s.cut(input)

		assert.Equal(t, output, result, testCase.name)
	}
}
