package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

func main() {
	input, _ := os.Open("./day14/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	occupied := map[Pos]bool{}
	max_depth := -1
	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)
		data := strings.Split(line, " -> ")

		var prev_x, prev_y int
		for i, pos := range data {
			d := strings.Split(pos, ",")
			x, _ := strconv.Atoi(d[0])
			y, _ := strconv.Atoi(d[1])

			if max_depth < y {
				max_depth = y
			}
			if i > 0 {
				dx := x - prev_x
				dy := y - prev_y
				step_x := dx / int(math.Max(1, math.Abs(float64(dx))))
				step_y := dy / int(math.Max(1, math.Abs(float64(dy))))

				for prev_x != x || prev_y != y {
					occupied[Pos{prev_x, prev_y}] = true
					prev_x += step_x
					prev_y += step_y
				}
				occupied[Pos{prev_x, prev_y}] = true
			} else {
				prev_x, prev_y = x, y
			}
		}
	}

	total := 0

	for !occupied[Pos{500, 0}] {
		sand_pos := Pos{500, 0}
		for sand_pos.y < max_depth+1 {

			if !occupied[Pos{sand_pos.x, sand_pos.y + 1}] {
				sand_pos.y++
			} else if !occupied[Pos{sand_pos.x - 1, sand_pos.y + 1}] {
				sand_pos.y++
				sand_pos.x--
			} else if !occupied[Pos{sand_pos.x + 1, sand_pos.y + 1}] {
				sand_pos.y++
				sand_pos.x++
			} else {
				break
			}
		}
		occupied[sand_pos] = true
		total++
	}

	fmt.Println(total)

}
