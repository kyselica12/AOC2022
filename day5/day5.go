package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	//Read input file
	input, _ := os.Open("./day5/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Split(bufio.ScanLines)

	var stacks [][]int32
	phase1 := true

	for sc.Scan() {

		line := sc.Text()

		if len(line) < 3 {

			for i, s := range stacks {
				//println(i, string(s))
				ReverseSlice(s)
				println(i, string(s))
			}

			phase1 = false
			//break

		}

		if phase1 {
			reading := false
			for i, char := range line {
				if char == '[' {
					reading = true
				} else if char == ']' {
					reading = false
				} else if reading {
					//fmt.Printf("%q, %d", char, i/4)
					stack_id := i / 4

					stacks = add_to_stacks(stacks, stack_id, char)
				}
			}
			//fmt.Print("\n")
		} else {
			var amount, from, to int
			_, err := fmt.Sscanf(line, "move %d from %d to %d", &amount, &from, &to)
			if err != nil {
				continue
			}
			//fmt.Println(amount, from, to, len(stacks[from-1]), len(stacks[to-1]), line)

			for i := 0; i < amount; i++ {
				n := len(stacks[from-1]) - 1
				x := stacks[from-1][n]
				stacks[to-1] = append(stacks[to-1], x)
				stacks[from-1] = stacks[from-1][:n]
			}

			//fmt.Println(amount, from, to, len(stacks[from-1]), len(stacks[to-1]))
			//break
		}
	}
	for _, s := range stacks {
		switch len(s) {
		case 0:
			print("")

		default:
			print(string(s[len(s)-1]))
		}
	}

}

func add_to_stacks(stacks [][]int32, stack_id int, char int32) [][]int32 {
	if len(stacks) < stack_id+1 {
		for len(stacks) <= stack_id {
			var tmp []int32
			stacks = append(stacks, tmp)
		}
	}
	stacks[stack_id] = append(stacks[stack_id], char)

	return stacks
}

func ReverseSlice[T comparable](s []T) {
	sort.SliceStable(s, func(i, j int) bool {
		return i > j
	})
}
