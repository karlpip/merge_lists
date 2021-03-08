package main

import (
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"sort"
	"os"
)

// Represents a interval with its lower and upper bounds
type interval struct {
	lowerBound uint64
	upperBound uint64
}

// Parses a string of intervals formatted as [(lower),(upper)] [(lower),(upper)] ...
// Bounds have to be unsigned to be parsed.
// Lower bound has to smaller than upper bound.
// Returns an array of interval objects, err is set if no parseable intervals have been found or if an interval with wrong bounds has been found
func parseIntervals(in string) (res []interval, err error) {
	re := regexp.MustCompile(`\[(\d+),\s?(\d+)\]`)
	rx_res := re.FindAllStringSubmatch(in, -1)
	if rx_res == nil {
		err = errors.New("no matches in input")
		return
	}

	// use the already known size for efficient allocation
	res = make([]interval, len(rx_res))
	for i, r := range rx_res {
		// trust the regex matching and ignore the error value
		lb, _ := strconv.ParseUint(r[1], 10, 64)
		ub, _ := strconv.ParseUint(r[2], 10, 64)
		if ub < lb {
			err = errors.New("upper bound is smaller than lower bound")
			return
		}
		res[i] = interval{lb, ub}
	}

	return
}

// Merges a list of intervals to a new list of merged overlapping intervals
// Complexity: O(n)
// Returns a the merged list
func mergeSortedIntervals(in []interval) []interval {
	var res []interval
	// avoid nil slice access in the first iteration
	res = append(res, in[0])

	for _, inter := range in {
		if inter.lowerBound > res[len(res)-1].upperBound {
			// inter cant be merged so it has to be in the result list
			res = append(res, inter)
		} else if inter.upperBound > res[len(res)-1].upperBound {
			// update the current intervals bounds / merge it
			res[len(res)-1].upperBound = inter.upperBound
		}
	}
	return res
}

func main() {
	input := flag.String("input", "[25,30] [2,19] [14, 23] [4,8]", "intervals to be merged")
	flag.Parse()

	intervals, err := parseIntervals(*input)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// needs go version >= 1.8
	// sort the slice to prepare for merging
	// Complexity: O(n * log(n)) src: https://golang.org/src/sort/sort.go line 227
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].lowerBound < intervals[j].lowerBound
	})

	mergedIntervals := mergeSortedIntervals(intervals)
	fmt.Println(mergedIntervals)

	os.Exit(0)
}
