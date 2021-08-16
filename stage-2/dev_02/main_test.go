package main

import "testing"

func Test_Unpack01(T *testing.T) {
	testCases := []struct {
		in       string
		out      string
		hasError bool
	}{
		{in: "вф4кк5", out: "вффффкккккк", hasError: false},
		{in: "a4bc2d5e", out: "aaaabccddddde", hasError: false},
		{in: "abcd", out: "abcd", hasError: false},
		{in: "45", out: "", hasError: true},
		{in: "a1", out: "a", hasError: false},
		{in: "a10", out: "aaaaaaaaaa", hasError: false},
		{in: "a20", out: "aaaaaaaaaaaaaaaaaaaa", hasError: false},
		{in: "a22", out: "aaaaaaaaaaaaaaaaaaaaaa", hasError: false},
		{in: "a100", out: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", hasError: false},
		{in: "asd0", out: "as", hasError: false},
		{in: `qwe\4\5`, out: "qwe45", hasError: false},
		{in: `qwe\45`, out: "qwe44444", hasError: false},
		{in: `qwe\\5`, out: `qwe\\\\\`, hasError: false},
	}
	for _, testCase := range testCases {
		v, err := UnPack(testCase.in)
		if err != nil && !testCase.hasError {
			T.Fatal("err: ", testCase.hasError, "!=", err)
		}
		if v != testCase.out {
			T.Fatal("val: ", testCase.out, "!=", v)
		}
	}
}
