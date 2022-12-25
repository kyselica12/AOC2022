package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y, z int
}

type SideCoordinates struct {
	ul, ur, dl, dr Point
}

type CubeSide struct {
	coords SideCoordinates
	data   [][]rune
}

func (c1 CubeSide) equals(c2 CubeSide) bool {
	return c1.coords == c2.coords
}

func main() {
	PATH := "./day22/test-input.txt"
	cube_size := 4

	input, _ := os.Open(PATH)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var lines [][]rune
	var cubes []CubeSide
	index := 0
	cube_index := 0
	cube_input_line := map[int]int{}
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

		if index%cube_size == 0 {
			for i := 0; i < len(striped_line); i += cube_size {
				var cube_data [][]rune
				for _, l := range lines[index-cube_size : index] {
					//fmt.Println(i, i+cube_size)
					cube_data = append(cube_data, l[i:i+cube_size])
				}

				y := (i + l1 - l2) / cube_size
				x := index/cube_size - 1
				cube := CubeSide{coords: SideCoordinates{ul: Point{x, y, 0}, ur: Point{x, y + 1, 0},
					dl: Point{x + 1, y, 0}, dr: Point{x + 1, y + 1, 0}},
					data: cube_data,
				}
				cubes = append(cubes, cube)

				cube_input_line[cube_index] = index - cube_size
				cube_index++
			}
		}
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
	cube := cubes[0]
	for len(rest) > 0 {
		n, _ := fmt.Sscanf(rest, "%d%c%s", &amount, &dir, &rest)
		//fmt.Printf("%d %d %d %c %c \n", x, y, amount, dirs[current_dir],  dir)

		if n >= 1 { // move in direction
			for i := 0; i < amount; i++ {
				x2 := x
				y2 := y

				switch dirs[current_dir] {
				case '>':
					y2++
				case '<':
					y2--
				case 'v':
					x2++
				case '^':
					x2--
				}

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

	}
	fmt.Println(1000*(x+1) + 4*(y+1+start_indices[x]) + current_dir)
}

func foldCube(cubes []CubeSide) {

	new_cube := []CubeSide{cubes[0]}

	c1 := new_cube[0]
	c1.coords.ur.z = 1

	for _, c2 := range cubes[1:] {

	}

}

func findSide(p1 Point, p2 Point, cube CubeSide, x int, y int, cubes []CubeSide) {

	for _, c := range cubes {
		if c.equals(cube) {
			continue
		}

		switch p1 {
		case c.coords.ul:
			if c.coords.ur == p2 { // go up
				return
			} else { // go left

			}
		case c.coords.ur:
		case c.coords.dl:
		case c.coords.dr:

		}

	}

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
