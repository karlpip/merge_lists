package main

import "testing"

func intervalsEqual(a []interval, b []interval) bool{
	if len(a) != len(b) {
		return false
	}
	for i, _ := range a {
		if a[i].upperBound != b[i].upperBound || a[i].lowerBound != b[i].lowerBound {
			return false
		}
	}

	return true
}

// Tests the parsing of intervals
func TestParsing(t *testing.T) {
	parseTestInputs := []string {
		"[1,2]", // One value
		"[1,2] [2,3]", // Normal values
		"[25,30] [2,19] [14, 23] [4,8]", // More values
		"[33, 22]", // Lower bound > upper bound
		"[-1, 2] [1a,2b]", // Negative values and hex values
		"[] [1,2", // Malformatted values
		"\n", // No values at all
	}

	parseTestResults := [][]interval {
		{interval{1, 2}},
		{interval{1, 2}, interval{2, 3}},
		{interval{25, 30}, interval{2, 19}, interval{14, 23}, interval{4, 8}},
		nil,
		nil,
		nil,
		nil,
	}

	for i, in := range parseTestInputs {
		res, err := parseIntervals(in)
		if err != nil {
			if parseTestResults[i] != nil {
				t.Errorf("parseIntervals failed unintentionally during test input %s error %s", in, err)
			}
			continue
		}
		if !intervalsEqual(res, parseTestResults[i]) {
			t.Errorf("parseIntervals returned unexpected list expected %v got %v ",  parseTestResults[i], res)
		}
	}
}

// Tests the merging of inputs
func TestMerging(t *testing.T) {
	mergeTestInputs := [][]interval {
		{interval{1, 1}}, // Single interval
		{interval{1, 2}, interval{2, 3}}, // 2 intervals 1 merge
		{interval{2, 19}, interval{4, 8}, interval{14, 23}, interval{25, 30}},
		{interval{800, 1300}, interval{1000, 1129}, interval{9000, 9001}, interval{18000, 60000}}, // 1 merge and 2 unmergeables
	}

	mergeTestResults := [][]interval {
		{interval{1, 1}},
		{interval{1, 3}},
		{interval{2,23}, interval{25, 30}},
		{interval{800, 1300}, interval{9000, 9001}, interval{18000, 60000}},
	}

	for i, in := range mergeTestInputs {
		res := mergeSortedIntervals(in)
		if !intervalsEqual(res, mergeTestResults[i]) {
			t.Errorf("parseIntervals returned unexpected list expected %v got %v ",  mergeTestResults[i], res)
		}
	}
}
