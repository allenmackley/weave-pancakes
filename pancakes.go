package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func all(vs []string, f func(string) bool) bool {
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

type testSuite struct {
	T         int
	TestCases []string
}

func newTestSuite(t int, testCases []string) (*testSuite, error) {
	if t < 1 {
		return nil, fmt.Errorf("there must be at least one test case")
	}
	if t > 100 {
		return nil, fmt.Errorf("too many test cases. Max 100")
	}
	r, _ := regexp.Compile("^[-+]+$")
	validSet := all(testCases, func(s string) bool {
		return r.MatchString(s)
	})
	if !validSet {
		return nil, fmt.Errorf("test case strings must include only '+' and '-' characters")
	}
	test := &testSuite{
		T:         t,
		TestCases: testCases,
	}
	return test, nil
}

type testCase struct {
	CaseNum  int
	NumFlips int
}

func newTestCase(caseNum int, numFlips int) *testCase {
	t := &testCase{
		CaseNum:  caseNum,
		NumFlips: numFlips,
	}
	return t
}

func (ts *testSuite) run() []*testCase {
	var wg sync.WaitGroup
	results := make([]*testCase, len(ts.TestCases))
	for i, v := range ts.TestCases {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			caseNum := i + 1
			numFlips := makeStackHappy(v)
			results[i] = newTestCase(caseNum, numFlips)
		}(i, v)
	}
	wg.Wait()
	return results
}

func main() {
	testCases := []string{
		"-",
		"-+",
		"+-",
		"+++",
		"--+-",
		// "-+-+--",
		// "---++--",
	}
	t := len(testCases)
	tests, err := newTestSuite(t, testCases)
	if err != nil {
		fmt.Println("Input error: ", err)
	} else {
		results := tests.run()
		for _, v := range results {
			caseNum := strconv.Itoa(v.CaseNum)
			numFlips := strconv.Itoa(v.NumFlips)
			fmt.Printf("Case #%v: %v\n", caseNum, numFlips)
		}
	}
}
