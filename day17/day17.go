package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Tuple struct {
	i, f int
}

func main() {
	input, _ := os.Open("./day17/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Scan()
	line := []rune(sc.Text())
	//fmt.Println(line)
	wind := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		wind[i] = int(line[i] - '=')
	}

	rocks := get_rocks()

	max_height := 0
	grid := [][]bool{
		[]bool{false, false, false, false, false, false, false},
	}

	n_rocks := 1000000000000
	floor := 0
	t := 0
	n_winds := len(wind)

	cache := map[string]Tuple{}

	for i := 0; i < n_rocks; i++ {

		r := rocks[i%5]
		x := 2
		y := max_height + 3
		stoped := false

		for !stoped {
			w := wind[t%n_winds]
			new_x := x + w
			t = (t + 1) % n_winds

			if new_x >= 0 && new_x+len(r[0]) <= 7 { //can move in wind direction
				if y > max_height { // its above other rocks
					x += w
				} else {
					if !colision(y, r, max_height, new_x, grid) {
						x += w
					}
				}
			}

			if y-1 > max_height {
				y--
			} else {
				if y > 0 && !colision(y-1, r, max_height, x, grid) {
					y--
				} else {
					// rock stoped
					stoped = true
					for j := y; j < minInt(y+len(r), max_height+1); j++ {
						rock_line := get_row(x, r[len(r)-1-(j-y)])

						for k := 0; k < 7; k++ {
							b := rock_line[k] || grid[j][k]
							grid[j][k] = b
						}
					}
					// add new rows to map
					for j := max_height + 1; j < y+len(r); j++ {
						//fmt.Println("_>>", j, len(r)-1-(j-y), max_height, y)
						rock_line := get_row(x, r[len(r)-1-(j-y)])
						grid = append(grid, rock_line)

					}

					if max_height < y+len(r) {
						grid = append(grid, make([]bool, 7))
					}

					new_max := maxInt(y+len(r), max_height)
					floor += new_max - max_height
					max_height = new_max
				}
			}
		}

		min := max_height
		for col := 0; col < 7; col++ {
			for j := max_height; j >= 0; j-- {
				if grid[j][col] {
					min = minInt(min, j)
					break
				}
			}
		}
		if min > 0 {
			grid = grid[min:]
			max_height -= min

			hash := strconv.Itoa(i%5) + strconv.Itoa(t)
			for j := 0; j < len(grid); j++ {
				for k := 0; k < len(grid[j]); k++ {
					if grid[j][k] {
						hash += "1"
					} else {
						hash += "0"
					}
				}
			}

			if _, ok := cache[hash]; ok {
				fmt.Println(";)")
				fmt.Println(floor, i)

				j := i
				prev := cache[hash]
				step_i := i - prev.i
				step_f := floor - prev.f
				height := floor

				for j+step_i < n_rocks {
					j += step_i
					height += step_f
				}

				i = j
				floor = height
				fmt.Println("New", i, floor)
			}

			cache[hash] = Tuple{i, floor}

		}

		if i == 62 {
			fmt.Println(floor)
			//break
		}
	}
	fmt.Println(max_height)
	print_grid(grid)
	fmt.Println(floor)
}

func print_grid(grid [][]bool) {

	for i := len(grid) - 1; i >= 0; i-- {
		tmp := ""
		for j := 0; j < 7; j++ {
			if grid[i][j] {
				tmp += "#"
			} else {
				tmp += "."
			}
		}
		fmt.Println(tmp)
	}
}

func get_row(x int, rock_line []bool) []bool {

	line := make([]bool, 7)
	for i := x; i < len(rock_line)+x; i++ {
		line[i] = rock_line[i-x]
	}
	return line
}

func colision(y int, r [][]bool, max_height int, x int, grid [][]bool) bool {
	conflict := false

	for j := y; j < minInt(y+len(r), max_height); j++ {
		rock_row := r[len(r)-1-(j-y)]

		for pos := x; pos < x+len(r[0]); pos++ {
			if grid[j][pos] && rock_row[pos-x] {
				conflict = true
				break
			}
		}
		if conflict {
			break
		}
	}
	return conflict
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func get_rocks() [][][]bool {
	shape1 := [][]bool{[]bool{true, true, true, true}}
	shape2 := [][]bool{
		[]bool{false, true, false},
		[]bool{true, true, true},
		[]bool{false, true, false},
	}
	shape3 := [][]bool{
		[]bool{false, false, true},
		[]bool{false, false, true},
		[]bool{true, true, true},
	}
	shape4 := [][]bool{
		[]bool{true},
		[]bool{true},
		[]bool{true},
		[]bool{true},
	}
	shape5 := [][]bool{
		[]bool{true, true},
		[]bool{true, true},
	}

	shapes := [][][]bool{shape1, shape2, shape3, shape4, shape5}
	return shapes
}
