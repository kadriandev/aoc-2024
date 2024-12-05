package main

import (
	"aoc-in-go/utils"
	"bufio"
	"github.com/jpillora/puzzler/harness/aoc"
	"strings"
)

func main() {
	aoc.Harness(run)
}

func isReportSafe(report []int) bool {
	last := report[0]
	increasing := report[0] < report[1]

	for i, n := range report {
		if i == 0 {
			continue
		}

		if increasing && (n < last || n-last < 1 || n-last > 3) {
			return false
		} else if !increasing && (n > last || last-n < 1 || last-n > 3) {
			return false
		}

		last = n
	}
	return true
}

func run(part2 bool, input string) any {
	r := strings.NewReader(input)
	scanner := bufio.NewScanner(r)

	safe_reports := 0

	if part2 {
		for scanner.Scan() {
			line := scanner.Text()
			report, err := utils.SliceAtoi(strings.Split(line, " "))
			utils.Check(err)

			for i := range report {
				subreport := make([]int, 0)

				subreport = append(subreport, report[:i]...)
				subreport = append(subreport, report[i+1:]...)

				if isReportSafe(subreport) {
					safe_reports++
					break
				}
			}
		}
		return safe_reports
	}

	for scanner.Scan() {
		line := scanner.Text()
		report, err := utils.SliceAtoi(strings.Split(line, " "))
		utils.Check(err)

		if isReportSafe(report) {
			safe_reports++
		}

	}
	return safe_reports
}
