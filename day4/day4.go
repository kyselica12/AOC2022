package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {

	var input, _ = os.ReadFile("./day4/input.txt")
	var score = 0
	for _, line := range strings.Split(string(input), "\n") {

		line = strings.TrimSpace(line)
		var elfs = strings.Split(line, ",")

		s1, e1 := get_indices(elfs[0])
		s2, e2 := get_indices(elfs[1])

		if s2 >= s1 && s2 <= e1 {
			score++
		} else if s1 >= s2 && s1 <= e2 {
			score++
		}
	}
	println(score)
}

func get_indices(elf string) (int, int) {
	var indices = strings.Split(elf, "-")
	var start, _ = strconv.Atoi(indices[0])
	var end, _ = strconv.Atoi(indices[1])
	return start, end
}
