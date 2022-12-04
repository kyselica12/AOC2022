package main

import (
	"fmt"
	"github.com/golang-collections/collections/set"
	"os"
	"strings"
)

func main() {

	part2()

}

func part1() {

	input, _ := os.ReadFile("./input.txt")

	var score = 0

	for _, sequence := range strings.Split(string(input), "\n") {

		var chars = []rune(sequence)

		var first_half = make([]rune, len(sequence))
		var used = make(map[rune]bool)

		for i, elem1 := range chars {

			if i < len(sequence)/2 {
				first_half = append(first_half, elem1)
			} else {
				for _, elem2 := range first_half {
					if elem2 == elem1 {
						if used[elem1] {
							continue
						}
						used[elem1] = true
						if elem1 > 'a'-1 {
							score += int(elem1-'a') + 1
							//fmt.Printf("%d, %q", int(elem1-'a')+1, elem1)
						} else {
							score += int(elem1-'A') + 27
							//fmt.Printf("%d, %q", int(elem1-'A')+27, elem1)
						}
					}
				}
			}

		}

	}
	fmt.Print(score)
}

func part2() {

	input, _ := os.ReadFile("./input.txt")

	var score = 0
	var count = make(map[rune]int)
	var res = set.New()

	for i, sequence := range strings.Split(string(input), "\n") {
		var chars = []rune(sequence)

		var actual = set.New()

		for _, elem := range chars {
			if elem >= 'a' && elem <= 'z' || elem >= 'A' && elem <= 'Z' {
				actual.Insert(elem)
			}
		}
		if i%3 == 0 {
			res = actual
		} else {
			res = res.Intersection(actual)
		}

		if i%3 == 2 {
			score += count_score(count)
			count = make(map[rune]int)
		}

	}
	score += count_score(count)
	fmt.Print(score)
}

func count_score(count map[rune]int) int {
	for elem, c := range count {
		if c == 3 {
			if elem > 'a'-1 {
				return int(elem-'a') + 1
			} else {
				return int(elem-'A') + 27
			}
		}
	}
	return 0
}
