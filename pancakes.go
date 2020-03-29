package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

func flipGroup(pancakes []string, pos int) {
	pSub := pancakes[:pos+1]
	for i, v := range pSub {
		if v == "-" {
			pancakes[i] = "+"
		} else {
			pancakes[i] = "-"
		}
	}
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
			flipGroup(pancakes, pos)
			numFlips++
		}
		pos--
	}
	return numFlips
}

func tests(t int, testCases []string) ([]string, error) {
	if t < 1 {
		return nil, fmt.Errorf("There must be at least one test case.")
	}
	if t > 100 {
		return nil, fmt.Errorf("Too many test cases. Max 100.")
	}
	r, _ := regexp.Compile("^[-+]+$")
	validSet := All(testCases, func(s string) bool {
		return r.MatchString(s)
	})
	if !validSet {
		return nil, fmt.Errorf("Test case strings must include only '+' and '-' characters.")
	}

	numFlips := []int{}
	results := []string{}
	for _, v := range testCases {
		numFlips = append(numFlips, makeStackHappy(v))
	}
	for i, v := range numFlips {
		results = append(results, "Case #"+strconv.Itoa(i)+": "+strconv.Itoa(v))
	}
	return results, nil
}

func main() {
	testCases := []string{
		"-",
		"-+",
		"+-",
		"+++",
		"--+-",
	}
	t := len(testCases)
	results, err := tests(t, testCases)
	if err != nil {
		fmt.Println("Input error: ", err)
	} else {
		fmt.Println("Results: ", results)
	}
}
