package main

import "testing"

func TestInverseCaptchaPt1(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"1122", 3},
		{"1111", 4},
		{"1234", 0},
		{"91212129", 9},
	}

	for _, testCase := range testCases {
		output := inverseCaptchaPt1(testCase.input)
		if output != testCase.expected {
			t.Errorf("Input %s gave incorrect ouput %d, expected %d.", testCase.input, output, testCase.expected)
		}
	}
}
