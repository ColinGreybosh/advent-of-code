package day04

import (
	"advent-of-code-2024/parsing"
)

type Answer struct {
	Part1 int64
}

func Solution(path string) (Answer, error) {
	wordSearch, err := parseFile(path)
	if err != nil {
		return Answer{}, err
	}
	part1, err := countWords(wordSearch, []rune{'X', 'M', 'A', 'S'})
	if err != nil {
		return Answer{}, err
	}
	return Answer{Part1: part1}, nil
}

type WordSearch struct {
	grid   map[int]map[int]rune
	width  int
	height int
}

func countWords(wordSearch WordSearch, word []rune) (int64, error) {
	var count int64
	for j := range wordSearch.height {
		for i := range wordSearch.width {
			count += countWordsInNeighborhood(wordSearch, word, i, j)
		}
	}
	return count, nil
}

type Direction struct {
	deltaX int
	deltaY int
}

var (
	NORTH      = Direction{deltaX: 0, deltaY: 1}
	SOUTH      = Direction{deltaX: 0, deltaY: -1}
	EAST       = Direction{deltaX: 1, deltaY: 0}
	WEST       = Direction{deltaX: -1, deltaY: 0}
	NORTH_EAST = Direction{deltaX: 1, deltaY: 1}
	NORTH_WEST = Direction{deltaX: -1, deltaY: 1}
	SOUTH_EAST = Direction{deltaX: 1, deltaY: -1}
	SOUTH_WEST = Direction{deltaX: -1, deltaY: -1}
)

var directions = []Direction{NORTH, SOUTH, EAST, WEST, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST}

func countWordsInNeighborhood(wordSearch WordSearch, word []rune, x int, y int) int64 {
	var count int64
	for _, direction := range directions {
		if isWordInDirection(wordSearch, word, x, y, direction) {
			count++
		}
	}
	return count
}

func isWordInDirection(wordSearch WordSearch, word []rune, x int, y int, direction Direction) bool {
	i := x
	j := y
	for _, targetLetter := range word {
		letter, found := wordSearch.grid[i][j]
		if !found || letter != targetLetter {
			return false
		}
		i += direction.deltaX
		j += direction.deltaY
	}
	return true
}

func parseFile(path string) (WordSearch, error) {
	lines, err := parsing.ReadLinesFromFile(path)
	if err != nil {
		return WordSearch{}, err
	}
	height := len(lines)
	width := len(lines[0])

	grid := make(map[int]map[int]rune)
	for j, line := range lines {
		for i, char := range line {
			if grid[i] == nil {
				grid[i] = make(map[int]rune)
			}
			grid[i][j] = char
		}
	}
	return WordSearch{grid, width, height}, nil
}
