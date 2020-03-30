package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func All(vs []string, f func(string) bool) bool {
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

type TestSuite struct {
	T         int
	TestCases []string
}

func NewTestSuite(t int, testCases []string) (*TestSuite, error) {
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
	test := &TestSuite{
		T:         t,
		TestCases: testCases,
	}
	return test, nil
}

type Test struct {
	CaseNum  int
	NumFlips int
}

func NewTest(caseNum int, numFlips int) *Test {
	t := &Test{
		CaseNum:  caseNum,
		NumFlips: numFlips,
	}
	return t
}

func (ts *TestSuite) Run() []*Test {
	var wg sync.WaitGroup
	results := make([]*Test, len(ts.TestCases))
	for i, v := range ts.TestCases {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
			caseNum := i + 1
			numFlips := makeStackHappy(v)
			results[i] = NewTest(caseNum, numFlips)
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
	tests, err := NewTestSuite(t, testCases)
	if err != nil {
		fmt.Println("Input error: ", err)
	} else {
		results := tests.Run()
		for _, v := range results {
			caseNum := strconv.Itoa(v.CaseNum)
			numFlips := strconv.Itoa(v.NumFlips)
			fmt.Printf("Case #%v: %v\n", caseNum, numFlips)
		}
	}
}
