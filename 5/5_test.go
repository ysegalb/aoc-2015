package main

import "testing"

func TestIsNice(t *testing.T) {
	testCases := []struct {
		input string
		nice  bool
	}{
		{"ugknbfddgicrmopn", true},
		{"aaa", true},
		{"jchzalrnumimnmhp", false},
		{"haegwjzuvuyypxyu", false},
		{"dvszwmarrgswjxmb", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			isNice := IsNice(tc.input)
			if isNice != tc.nice {
				t.Errorf("Result for (%q): %t; want %t", tc.input, isNice, tc.nice)
			}
		})
	}
}

func TestIsNice2(t *testing.T) {
	testCases := []struct {
		input string
		nice  bool
	}{
		{"qjhvhtzxzqqjkmpb", true},
		{"xxyxx", true},
		{"uurcxstgmygtbstg", false},
		{"ieodomkazucvgmuy", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			isNice := IsNice2(tc.input)
			if isNice != tc.nice {
				t.Errorf("Result for (%q): %t; want %t", tc.input, isNice, tc.nice)
			}
		})
	}
}
