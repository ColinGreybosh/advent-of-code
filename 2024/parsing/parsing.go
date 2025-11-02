package parsing

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadLinesFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

func ParseLine(line string) ([]int64, error) {
	elements := strings.Split(line, " ")
	var filtered_elements = []int64{}
	for _, element := range elements {
		if element == "" {
			continue
		}
		value, err := strconv.ParseInt(element, 10, 64)
		if err != nil {
			return nil, err
		}
		filtered_elements = append(filtered_elements, value)
	}
	return filtered_elements, nil
}
