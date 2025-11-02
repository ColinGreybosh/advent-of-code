package Day02

import (
	"advent-of-code-2024/math"
	"advent-of-code-2024/parsing"
	"iter"
)

type Answer struct {
	Part1 int64
	Part2 int64
}

func Solution(path string) (Answer, error) {
	reports, err := parseFile(path)
	if err != nil {
		return Answer{}, nil
	}
	var safeReportCount int64 = 0
	var safeReportCountWithDampening int64 = 0
	for _, report := range reports {
		if isReportSafe(report, false) {
			safeReportCount++
		}
		if isReportSafe(report, true) {
			safeReportCountWithDampening++
		}
	}
	return Answer{
		Part1: safeReportCount,
		Part2: safeReportCountWithDampening,
	}, nil
}

type LevelDirection string

const (
	INCREASING LevelDirection = "INCREASING"
	DECREASING LevelDirection = "DECREASING"
	UNKNOWN    LevelDirection = "UNKNOWN"
)

func isReportSafe(report []int64, shouldDampen bool) bool {
	currentDirection := UNKNOWN
	for i := 1; i < len(report); i++ {
		previousLevel := report[i-1]
		currentLevel := report[i]
		difference := math.Abs(previousLevel - currentLevel)
		if !(1 <= difference && difference <= 3) {
			// Unsafe difference in adjacent levels
			if shouldDampen {
				return isDampenedReportSafe(report)
			}
			return false
		}
		newDirection := UNKNOWN
		if currentLevel > previousLevel {
			newDirection = INCREASING
		} else {
			newDirection = DECREASING
		}
		if currentDirection != UNKNOWN && newDirection != currentDirection {
			// Unsafe change in direction
			if shouldDampen {
				return isDampenedReportSafe(report)
			}
			return false
		}
		if currentDirection == UNKNOWN {
			currentDirection = newDirection
		}
	}
	return true
}

func isDampenedReportSafe(report []int64) bool {
	dampenedReports := getDampenedReports(report)
	for dampenedReport := range dampenedReports {
		if isReportSafe(dampenedReport, false) {
			return true
		}
	}
	return false
}

func getDampenedReports(report []int64) iter.Seq[[]int64] {
	return func(yield func([]int64) bool) {
		for i := range len(report) {
			var dampenedReport []int64
			dampenedReport = append(dampenedReport, report[:i]...)
			dampenedReport = append(dampenedReport, report[i+1:]...)
			if !yield(dampenedReport) {
				return
			}
		}
	}
}

func parseFile(path string) ([][]int64, error) {
	lines, err := parsing.ReadLinesFromFile(path)
	if err != nil {
		return nil, err
	}
	reports := [][]int64{}
	for _, lineToParse := range lines {
		parsedLine, err := parsing.ParseLine(lineToParse)
		if err != nil {
			return nil, err
		}
		reports = append(reports, parsedLine)
	}
	return reports, nil
}
