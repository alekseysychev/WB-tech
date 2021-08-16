package main

import (
	"fmt"
	"testing"
)

func Test_sss(t *testing.T) {
	for _, testCase := range []struct {
		name   string
		word   string
		flags  flagsStrusct
		input  string
		output string
	}{
		{
			name:   `grep -A 10 "F.*" input.txt > grep-A.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{A: 10},
			input:  "input.txt",
			output: "grep-A.txt",
		},
		{
			name:   `grep -A 2 -B 1 "F.*" input.txt > grep-A2-B1.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{A: 2, B: 1},
			input:  "input.txt",
			output: "grep-A2-B1.txt",
		},
		{
			name:   `grep -B 10 "F.*" input.txt > grep-B.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{B: 10},
			input:  "input.txt",
			output: "grep-B.txt",
		},
		{
			name:   `grep -C 2 "F.*" input.txt > grep-C.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{C: 2},
			input:  "input.txt",
			output: "grep-C.txt",
		},
		{
			name:   `grep -c "F.*" input.txt > grep-c.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{c: true},
			input:  "input.txt",
			output: "grep-c.txt",
		},
		{
			name:   `grep -i "F.*" input.txt > grep-i.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{i: true},
			input:  "input.txt",
			output: "grep-i.txt",
		},
		{
			name:   `grep -v "F.*" input.txt > grep-v.txt`,
			word:   `F.*`,
			flags:  flagsStrusct{v: true},
			input:  "input.txt",
			output: "grep-v.txt",
		},
		{
			name:   `grep -F "f.*" input.txt > grep-F.txt`,
			word:   `f.*`,
			flags:  flagsStrusct{F: true},
			input:  "input.txt",
			output: "grep-F.txt",
		},
		{
			name:   `grep -n "f.*" input.txt > grep-n.txt`,
			word:   `f.*`,
			flags:  flagsStrusct{n: true},
			input:  "input.txt",
			output: "grep-n.txt",
		},
	} {

		input, _ := getRowsFromFile(testCase.input)
		output, _ := getRowsFromFile(testCase.output)

		s := grepStruct{
			flags: testCase.flags,
			word:  testCase.word,
		}
		result := s.grep(input)

		compare := compareRows(output, result)

		if compare != nil {
			fmt.Println("input:::", len(input))
			// for i := 0; i < len(input); i++ {
			// 	fmt.Println(input[i])
			// }
			fmt.Println("output:::", len(output))
			// for i := 0; i < len(output); i++ {
			// fmt.Println(output[i])
			// }
			fmt.Println("result:::", len(result))
			t.Error(testCase.name, *compare)
			// вывод результирующих строк
			for i := 0; i < len(result); i++ {
				fmt.Println(result[i])
			}
			break
		}

	}
}
