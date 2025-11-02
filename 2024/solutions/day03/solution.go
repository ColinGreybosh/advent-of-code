package Day03

import (
	"advent-of-code-2024/parsing"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Answer struct {
	Part1 int64
	Part2 int64
}

func Solution(path string) (Answer, error) {
	part1Multiplications, err := extractMultiplications(path, false)
	if err != nil {
		return Answer{}, nil
	}
	part1 := calculateSum(part1Multiplications)

	part2Multiplications, err := extractMultiplications(path, true)
	if err != nil {
		return Answer{}, nil
	}
	part2 := calculateSum(part2Multiplications)
	return Answer{Part1: part1, Part2: part2}, nil
}

type Multiplication struct {
	left  int64
	right int64
}

func calculateSum(multiplications []Multiplication) int64 {
	var sum int64
	for _, multiplication := range multiplications {
		sum += multiplication.left * multiplication.right
	}
	return sum
}

func extractMultiplications(path string, isPartTwo bool) ([]Multiplication, error) {
	lines, err := parsing.ReadLinesFromFile(path)
	if err != nil {
		return nil, err
	}
	re, err := regexp.Compile(`(?P<mul>mul\((?P<left>\d{1,3}),(?P<right>\d{1,3})\))|(?P<do>do\(\))|(?P<dont>don't\(\))`)
	if err != nil {
		return nil, err
	}
	subexpnames := re.SubexpNames()
	leftIndex := slices.Index(subexpnames, "left")
	rightIndex := slices.Index(subexpnames, "right")
	var multiplications []Multiplication
	isOn := true
	for _, line := range lines {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if isOn && isMultiplication(match) {
				multiplication, err := parseMultiplication(match, leftIndex, rightIndex)
				if err != nil {
					return nil, err
				}
				multiplications = append(multiplications, multiplication)
			}
			if isStart(match) {
				isOn = true
			}
			if isPartTwo && isStop(match) {
				isOn = false
			}
		}
	}
	return multiplications, nil
}

func isMultiplication(match []string) bool {
	_, found := strings.CutPrefix(match[0], "mul(")
	return found
}

func isStart(match []string) bool {
	return match[0] == "do()"
}

func isStop(match []string) bool {
	return match[0] == "don't()"
}

func parseMultiplication(match []string, leftIndex int, rightIndex int) (Multiplication, error) {
	left, err := strconv.ParseInt(match[leftIndex], 10, 64)
	if err != nil {
		return Multiplication{}, err
	}
	right, err := strconv.ParseInt(match[rightIndex], 10, 64)
	if err != nil {
		return Multiplication{}, err
	}
	return Multiplication{left, right}, nil
}
