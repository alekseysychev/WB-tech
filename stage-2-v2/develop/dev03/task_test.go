package main

import (
	"fmt"
	"testing"
)

func Test_sss(t *testing.T) {
	for _, testCase := range []struct {
		name   string
		flags  map[string]interface{}
		input  string
		output string
	}{
		{
			name:   "LC_ALL=C sort allkeys.txt > allkeys-sorted.txt",
			flags:  map[string]interface{}{},
			input:  "allkeys.txt",
			output: "allkeys-sorted.txt",
		},
		{
			name:   "LC_ALL=C sort -r allkeys.txt > allkeys-sorted-r.txt",
			flags:  map[string]interface{}{"r": true},
			input:  "allkeys.txt",
			output: "allkeys-sorted-r.txt",
		},
		{
			name:   "LC_ALL=C sort -k2 -n allkeys.txt > allkeys-sorted-k2-n.txt",
			flags:  map[string]interface{}{"n": true, "k": 2},
			input:  "allkeys.txt",
			output: "allkeys-sorted-k2-n.txt",
		},
		{
			name:   "LC_ALL=C sort -u allkeys.txt > allkeys-sorted-u.txt",
			flags:  map[string]interface{}{"u": true},
			input:  "allkeys.txt",
			output: "allkeys-sorted-u.txt",
		},
		{
			name:   "LC_ALL=C sort -k3 -M allkeys.txt > allkeys-sorted-k3-M.txt",
			flags:  map[string]interface{}{"M": true, "k": 3},
			input:  "allkeys.txt",
			output: "allkeys-sorted-k3-M.txt",
		},
		{
			name:   "LC_ALL=C sort -b allkeys.txt > allkeys-sorted-b.txt",
			flags:  map[string]interface{}{"b": true},
			input:  "allkeys.txt",
			output: "allkeys-sorted-b.txt",
		},
		{
			name:   "LC_ALL=C sort -k4 -h allkeys.txt > allkeys-sorted-k4-h.txt",
			flags:  map[string]interface{}{"h": true, "k": 4},
			input:  "allkeys.txt",
			output: "allkeys-sorted-k4-h.txt",
		},
	} {

		input, _ := getRowsFromFile(testCase.input)
		output, _ := getRowsFromFile(testCase.output)

		s := sortStruct{
			flags: testCase.flags,
		}
		result := s.sort(input)

		compare := compareRows(output, result)

		if compare != nil {
			fmt.Println("input:")
			for i := 0; i < len(input); i++ {
				fmt.Println(input[i])
			}
			fmt.Println("output:")
			for i := 0; i < len(output); i++ {
				fmt.Println(output[i])
			}
			fmt.Println("result")
			t.Error(testCase.name, *compare)
			// вывод результирующих строк
			for i := 0; i < len(result); i++ {
				fmt.Println(result[i])
			}
			break
		}

	}
}
