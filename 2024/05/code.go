package main

import (
	"aoc-in-go/utils"
	"bufio"
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func isUpdateValid(update []string, rules [][]string) bool {
	new_update := []string{}
	for _, num := range update {
		if len(new_update) > 0 {
			applicable_rules := utils.Filter(rules, func(s []string) bool { return s[0] == num && slices.Contains(update, s[1]) })
			for _, rule := range applicable_rules {
				if slices.Contains(new_update, rule[1]) {
					return false
				}
			}
		}
		new_update = append(new_update, num)
	}
	return true
}

func makeUpdateValid(update []string, rules [][]string) []string {
	new_update := slices.Clone(update)

	applicable_rules := utils.Filter(rules, func(s []string) bool {
		return slices.Contains(update, s[0]) && slices.Contains(update, s[1])
	})

	slices.SortFunc(new_update, func(a, b string) int {
		idx := slices.IndexFunc(applicable_rules, func(s []string) bool {
			return slices.Contains(s, a) && slices.Contains(s, b)
		})
		idx_a := slices.Index(applicable_rules[idx], a)
		if idx_a == 0 {
			return -1
		}
		return 1
	})

	return new_update
}

func run(part2 bool, input string) any {
	middles := []int{}

	rules := [][]string{}
	updates := [][]string{}

	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		rules = append(rules, strings.Split(line, "|"))
	}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		updates = append(updates, strings.Split(line, ","))
	}

	if part2 {
		for _, update := range updates {
			if !isUpdateValid(update, rules) {
        fixed := makeUpdateValid(update, rules)
				update_len := len(fixed)
				middles = append(middles, utils.Str_to_int(fixed[(update_len/2)]))
			}
		}

		return utils.Sum(middles)
	}

	for _, update := range updates {
		if isUpdateValid(update, rules) {
			update_len := len(update)
			middles = append(middles, utils.Str_to_int(update[(update_len/2)]))
		}
	}

	return utils.Sum(middles)
}
