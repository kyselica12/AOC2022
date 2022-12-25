package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("./day22/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var lines [][]rune
	index := 0
	start_indices := map[int]int{}

	for sc.Scan() {
		line := sc.Text()
		striped_line := strings.TrimSpace(line)

		l1 := len(line)
		l2 := len(striped_line)

		if l2 == 0 {
			break
		}

		start_indices[index] = l1 - l2
		lines = append(lines, []rune(striped_line))

		index++
	}

	sc.Scan()
	line := strings.TrimSpace(sc.Text())
	var amount int
	var dir rune
	rest := line

	dirs := map[int]rune{0: '>', 1: 'v', 2: '<', 3: '^'}
	current_dir := 0
	x := 0
	y := 0

	//print_board(x, y, lines, start_indices)

	for len(rest) > 0 {
		n, _ := fmt.Sscanf(rest, "%d%c%s", &amount, &dir, &rest)
		fmt.Printf("%d %d %d %c %c \n", x, y, amount, dirs[current_dir], dir)
		//fmt.Println(n, x, y, amount, dir)
		if x == 148 && y == 50 {
			fmt.Println("HERE")
		}

		if n >= 1 { // move in direction
			for i := 0; i < amount; i++ {
				x2 := x
				y2 := y

				switch dirs[current_dir] {
				case '>':
					y2 = (y + 1) % len(lines[x])
				case '<':
					y2 = (y - 1 + len(lines[x])) % len(lines[x])
				case 'v':
					x2++
					y2 = y2 + start_indices[x2-1] - start_indices[x2]
					if x2 >= len(lines) || len(lines[x2]) <= y2 || y2 < 0 { // roll
						x2--
						y2 = y2 + start_indices[x2+1] - start_indices[x2]
						for x2 >= 0 && len(lines[x2]) > y2 && y2 >= 0 {
							x2--
							y2 = y2 + start_indices[x2+1] - start_indices[x2]
						}
						x2++
						y2 = y2 + start_indices[x2-1] - start_indices[x2]
					}
				case '^':
					x2--
					y2 = y2 + start_indices[x2+1] - start_indices[x2]
					if x2 < 0 || len(lines[x2]) <= y2 || y2 < 0 { // roll
						x2++
						y2 = y2 + start_indices[x2-1] - start_indices[x2]
						for x2 < len(lines) && len(lines[x2]) > y2 && y2 >= 0 {
							x2++
							y2 = y2 + start_indices[x2-1] - start_indices[x2]
						}
						x2--
						y2 = y2 + start_indices[x2+1] - start_indices[x2]
					}
				}
				//fmt.Println(x2,y2)
				if lines[x2][y2] == '#' {
					break
				}
				x = x2
				y = y2
			}
		}
		if n >= 2 {
			switch dir {
			case 'R':
				current_dir = (current_dir + 1) % 4
			case 'L':
				current_dir = (current_dir - 1 + 4) % 4
			}
		}
		if n == 1 {
			break
		}
		//print_board(x, y, lines, start_indices)

	}
	fmt.Println(1000*(x+1) + 4*(y+1+start_indices[x]) + current_dir)
}

func print_board(x, y int, lines [][]rune, start_indices map[int]int) {

	for i := 0; i < len(lines); i++ {

		for j := 0; j < start_indices[i]; j++ {
			fmt.Print(" ")
		}
		for y0 := 0; y0 < len(lines[i]); y0++ {
			if i == x && y0 == y {
				fmt.Print("X")
			} else {
				fmt.Printf("%c", lines[i][y0])
			}
		}
		fmt.Print("\n")
	}
	fmt.Println("\n")
}
