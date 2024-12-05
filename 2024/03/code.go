package main

import (
	"aoc-in-go/utils"
	"regexp"
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
	result := 0

	if part2 {
		enabled := true
		r, _ := regexp.Compile("(do\\(\\))|(don't\\(\\))|(mul\\([0-9]+,[0-9]+\\))")

		muls := r.FindAllString(input, 1000)
		for _, exp := range muls {
			if strings.HasPrefix(exp, "don't") {
				enabled = false
			} else if strings.HasPrefix(exp, "do") {
				enabled = true
			} else if strings.HasPrefix(exp, "mul") && enabled {
				nums := exp[4:]
				nums = nums[:len(nums)-1]
				numArr := strings.Split(nums, ",")
				result += utils.Str_to_int(numArr[0]) * utils.Str_to_int(numArr[1])
			}
		}
		return result
	}

	r, _ := regexp.Compile("mul\\([0-9]+,[0-9]+\\)")
	muls := r.FindAllString(input, 1000)
	for _, exp := range muls {
		nums := exp[4:]
		nums = nums[:len(nums)-1]
		numArr := strings.Split(nums, ",")
		result += utils.Str_to_int(numArr[0]) * utils.Str_to_int(numArr[1])
	}

	return result
}
