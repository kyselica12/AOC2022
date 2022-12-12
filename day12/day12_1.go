// convert to Python3

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func main() {
	file_name := "./day12/input.txt"
	world := read_file(file_name)

	var start, end Pos

	for i, line := range world {
		for j, x := range line {
			if x == 'S' {
				start = Pos{x: i, y: j}
				world[i][j] = 'a'
			}
			if x == 'E' {
				end = Pos{x: i, y: j}
				world[i][j] = 'z'
			}
		}
	}
	println(len(world), len(world[0]))
	queue := [][]Pos{[]Pos{start}}
	var winner []Pos

	visited := map[Pos]bool{}
	visited[start] = true
	fmt.Println(start, end)
	total := 0
	for len(queue) > 0 {

		route := queue[0]
		last := route[len(route)-1]
		queue = queue[1:]
		if last == end {
			winner = route
			break
		}

		directions := []Pos{Pos{x: last.x - 1, y: last.y}, Pos{x: last.x + 1, y: last.y},
			Pos{x: last.x, y: last.y - 1}, Pos{x: last.x, y: last.y + 1}}
		for _, to := range directions {
			if can_move(world, last, to) {
				if !visited[to] {
					var new_route []Pos
					new_route = append(new_route, route...)
					new_route = append(new_route, to)

					queue = append(queue, new_route)
					visited[to] = true
					total += 1
				}
			}
		}
	}

	//fmt.Println(winner, visited[end])
	fmt.Println(len(winner)-1, total, len(world)*len(world[0]))
}

func can_move(world [][]rune, from Pos, to Pos) bool {
	return to.x >= 0 && to.x < len(world) && to.y >= 0 && to.y < len(world[0]) &&
		world[to.x][to.y]-world[from.x][from.y] <= 1
}

func read_file(file_name string) [][]rune {
	input, _ := os.Open(file_name)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var world [][]rune

	for sc.Scan() {

		line := sc.Text()
		row := []rune(strings.TrimSpace(line))

		world = append(world, row)
	}
	return world
}
