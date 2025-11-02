package Day01

import (
	"advent-of-code-2024/math"
	"advent-of-code-2024/parsing"
	"fmt"
	"slices"
)

type Answer struct {
	Part1 int64
	Part2 int64
}

func Solution(path string) (Answer, error) {
	lists, err := parseFile(path)
	if err != nil {
		return Answer{}, err
	}

	distance, err := calculateDistance(lists)
	if err != nil {
		return Answer{}, nil
	}
	similarity, err := calculateSimilarity(lists)
	if err != nil {
		return Answer{}, nil
	}
	return Answer{Part1: distance, Part2: similarity}, nil
}

type Lists struct {
	Left  []int64
	Right []int64
}

func calculateDistance(lists Lists) (int64, error) {
	slices.Sort(lists.Left)
	slices.Sort(lists.Right)

	var distance int64
	for index := range lists.Left {
		distance += math.Abs(lists.Left[index] - lists.Right[index])
	}
	return distance, nil
}

func calculateSimilarity(lists Lists) (int64, error) {
	counts, err := countNumbers(lists.Right)
	if err != nil {
		return 0, err
	}
	var similarity int64
	for _, number := range lists.Left {
		similarity += number * counts[number]
	}
	return similarity, nil
}

func countNumbers(list []int64) (map[int64]int64, error) {
	var counts = make(map[int64]int64)
	for _, number := range list {
		counts[number] = counts[number] + 1
	}
	return counts, nil
}

func parseFile(path string) (Lists, error) {
	lines, err := parsing.ReadLinesFromFile(path)
	if err != nil {
		return Lists{}, err
	}
	left, right := []int64{}, []int64{}

	for _, lineToParse := range lines {
		parsedLine, err := parsing.ParseLine(lineToParse)
		if err != nil {
			return Lists{}, err
		}
		if len(parsedLine) != 2 {
			return Lists{}, fmt.Errorf("expected line with 2 elements but saw %d", len(parsedLine))
		}
		left = append(left, parsedLine[0])
		right = append(right, parsedLine[1])
	}
	return Lists{left, right}, nil
}
