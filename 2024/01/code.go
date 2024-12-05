package main

import (
	"aoc-in-go/utils"
	"bufio"
	"math"
	"sort"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

	var list1 []int
	var list2 []int

	for scanner.Scan() {
		line := scanner.Text()
		tokens := strings.Split(line, " ")

		list1 = append(list1, utils.Str_to_int(tokens[0]))
		list2 = append(list2, utils.Str_to_int(tokens[3]))
	}

	if part2 {
		similarity := 0

		for _, n := range list1 {
			filtered := []int{}

			for _, n2 := range list2 {
				if n == n2 {
					filtered = append(filtered, n2)
				}
			}
			similarity += n * len(filtered)
		}

		return similarity
	}

	distance := 0

	sort.Ints(list1)
	sort.Ints(list2)

	for i := range list1 {
		d := list2[i] - list1[i]
		distance += int(math.Abs(float64(d)))
	}

	// solve part 1 here
	return distance
}
