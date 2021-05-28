package main

import (
	//"fmt"
	"testing"
)

func TestConvertKeyModeIntToKey(t *testing.T) {
	var tests = []struct{
		key int
		mode int
		expected string
	} {
		{0,0, "Cm"},
		{4,1, "E"},
		{12,1, ""},
		{11,1, "B"},
	}

	for _, test := range tests {
		if output := convertKeyModeIntToKey(test.key,test.mode); output != test.expected {
			t.Errorf("Test failed: (%v, %v) inputted, %v expected, received: %v",test.key, test.mode, test.expected, output)
		}
	}
}