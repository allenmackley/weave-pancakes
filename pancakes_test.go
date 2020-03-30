package main

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type testCase struct {
	value    string
	expected int
}

func all(vs []testCase, f func(testCase) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func flipGroup(pancakes []string, flipToIndex int) {
	pSub := reverse(pancakes[:flipToIndex+1])
	for i, v := range pSub {
		if v == "-" {
			pancakes[i] = "+"
		} else {
			pancakes[i] = "-"
		}
	}
}

func flipGroupsFrom(pancakes []string, flipFromIndex int) int {
	var numFlips int
	flipToIndex := -1
	if string(pancakes[0]) == "+" {
		for i := 0; i < len(pancakes); i++ {
			if string(pancakes[i]) == "+" {
				flipToIndex = i
			} else {
				break
			}
		}
	}
	if flipToIndex > -1 {
		flipGroup(pancakes, flipToIndex)
		numFlips++
	}
	flipGroup(pancakes, flipFromIndex)
	numFlips++
	return numFlips
}

func makeStackHappy(s string) int {
	var numFlips int
	var needToFlip bool
	pancakes := strings.Split(s, "")
	l := len(pancakes)
	pos := l - 1
	for pos >= 0 && numFlips <= l {
		needToFlip = string(pancakes[pos]) == "-"
		if needToFlip {
			numFlips += flipGroupsFrom(pancakes, pos)
		}
		pos--
	}
	return numFlips
}

func TestMany(t *testing.T) {
	testCases := []testCase{
		{"-", 1},
		{"-+", 1},
		{"+-", 2},
		{"+++", 0},
		{"--+-", 3},
		// Uncomment these two lines for a more comprehensive test. These tests ensure that the pancakes are flipped in a group from the top and not one at a time from the middle.
		// {"-+-+--", 5},
		// {"---++--", 3},
	}
	numTests := len(testCases)
	if numTests < 1 {
		t.Fatalf("there must be at least one test case")
	}
	if numTests > 100 {
		t.Fatalf("too many test cases, max 100")
	}
	r, _ := regexp.Compile("^[-+]+$")
	validSet := all(testCases, func(tc testCase) bool {
		return r.MatchString(tc.value)
	})
	if !validSet {
		t.Fatalf("test case strings must include only '+' and '-' characters")
	}
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("%s", tc.value), func(t *testing.T) {
			numFlips := makeStackHappy(tc.value)
			if numFlips != tc.expected {
				t.Fatalf("expected %d flips, got %d", tc.expected, numFlips)
			}
			fmt.Printf("Case #%v: %v\n", i+1, numFlips)
		})
	}
}

func main() {}
