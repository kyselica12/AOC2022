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

type State struct {
	p Pos
	t int
}

type Storm struct {
	dx, dy int
	p      Pos
}

func main() {
	PATH := "./day24/input.txt"
	walls, storms, start, end, height, width := read_file_day24(PATH)

	T := 0
	queue := map[int]map[Pos]bool{}
	queue[0] = map[Pos]bool{start: true}
	found := false

	checkpoint := 0
	checkpoints := []Pos{end, start, end}

	for !found {
		storms_pos := update_storms(storms, walls, height, width)
		queue[T+1] = map[Pos]bool{}
		fmt.Println(len(queue[T]))
		for p, ok := range queue[T] {
			if !ok {
				continue
			}

			if p == checkpoints[checkpoint] {
				queue[T+1] = map[Pos]bool{}
			}

			moves := []Pos{{p.x - 1, p.y}, {p.x + 1, p.y}, {p.x, p.y}, {p.x, p.y + 1}, {p.x, p.y - 1}}
			for _, move := range moves {
				if move.x < 0 || move.y < 0 || move.x >= height || move.y >= width || storms_pos[move] || walls[move] {
					continue
				}
				queue[T+1][move] = true
			}
			if p == checkpoints[checkpoint] {
				fmt.Print("END", p, T)
				checkpoint++
				if checkpoint == len(checkpoints) {
					found = true
				}

				break
			}

		}
		T++
	}
}

func update_storms(storms []Storm, walls map[Pos]bool, height int, width int) map[Pos]bool {
	storms_pos := map[Pos]bool{}
	for i, storm := range storms {
		new_pos := Pos{storm.p.x + storm.dx, storm.p.y + storm.dy}

		if walls[new_pos] {
			if storm.dx == 1 {
				new_pos.x = 1
			} else if storm.dx == -1 {
				new_pos.x = height - 2
			}

			if storm.dy == 1 {
				new_pos.y = 1
			} else if storm.dy == -1 {
				new_pos.y = width - 2
			}
		}
		storms[i].p = new_pos
		storms_pos[new_pos] = true
	}
	return storms_pos
}

func print_board(walls map[Pos]bool, storms []Storm, p Pos, height, width int) {
	fmt.Println("-----------------")
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			pos := Pos{i, j}
			direction := ""
			for _, s := range storms {
				if s.p == pos {
					if s.dy == 1 {
						direction += ">"
					} else if s.dy == -1 {
						direction += "<"
					} else if s.dx == 1 {
						direction += "v"
					} else if s.dx == -1 {
						direction += "^"
					}
				}
			}
			if len(direction) > 0 {
				if len(direction) < 2 {
					fmt.Print(direction)
				} else {
					fmt.Print(len(direction))
				}
			} else if walls[pos] {
				fmt.Print("#")
			} else if p == pos {
				fmt.Print("+")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("-----------------\n")

}

func read_file_day24(PATH string) (map[Pos]bool, []Storm, Pos, Pos, int, int) {
	input, _ := os.Open(PATH)
	defer input.Close()
	sc := bufio.NewScanner(input)

	walls := map[Pos]bool{}
	var storms []Storm
	row := 0

	width := 0
	for sc.Scan() {
		line := sc.Text()
		striped_line := strings.TrimSpace(line)

		for col, c := range striped_line {
			switch c {
			case '#':
				walls[Pos{row, col}] = true
			case '>':
				storms = append(storms, Storm{0, 1, Pos{row, col}})
			case '<':
				storms = append(storms, Storm{0, -1, Pos{row, col}})
			case '^':
				storms = append(storms, Storm{-1, 0, Pos{row, col}})
			case 'v':
				storms = append(storms, Storm{1, 0, Pos{row, col}})
			}
		}
		width = len(striped_line)
		row++
	}
	height := row
	start := Pos{0, 1}
	end := Pos{height - 1, width - 2}

	return walls, storms, start, end, row, width
}
