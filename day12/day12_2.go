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
	queue := []Pos{start}
	visited := map[Pos]int{}

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
			if world[i][j] == 'a' {
				p := Pos{i, j}
				queue = append(queue, p)
				visited[p] = 1
			}
		}
	}
	println(len(world), len(world[0]))

	for len(queue) > 0 {

		last := queue[0]
		queue = queue[1:]
		if last == end {
			fmt.Println("winner: ", visited[end]-1)
			break
		}

		directions := []Pos{Pos{x: last.x - 1, y: last.y}, Pos{x: last.x + 1, y: last.y},
			Pos{x: last.x, y: last.y - 1}, Pos{x: last.x, y: last.y + 1}}
		for _, to := range directions {
			if can_move(world, last, to) {
				if visited[to] == 0 {
					queue = append(queue, to)
					visited[to] = visited[last] + 1
				}
			}
		}
	}

	fmt.Println(visited[end])
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
