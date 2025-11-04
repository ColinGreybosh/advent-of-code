package main

import (
	Day01 "advent-of-code-2024/solutions/day01"
	Day02 "advent-of-code-2024/solutions/day02"
	Day03 "advent-of-code-2024/solutions/day03"
	Day04 "advent-of-code-2024/solutions/day04"
	"fmt"
)

func main() {
	fmt.Println("====== DAY 01 ======")
	answer1, err := Day01.Solution("inputs/day01.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer1)
	}

	fmt.Println("\n====== DAY 02 ======")
	answer2, err := Day02.Solution("inputs/day02.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer2)
	}

	fmt.Println("\n====== DAY 03 ======")
	answer3, err := Day03.Solution("inputs/day03.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer3)
	}

	fmt.Println("\n====== DAY 04 ======")
	answer4, err := Day04.Solution("inputs/day04.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(answer4)
	}
}
