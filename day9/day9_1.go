package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Index struct {
	x, y int
}

func main() {
	input, _ := os.Open("./day9/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	visited := map[Index]bool{}
	N := 10

	var snake []Index
	for i := 0; i < N; i++ {
		snake = append(snake, Index{0, 0})
	}

	visited[snake[N-1]] = true

	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		var direction string
		var steps int

		fmt.Sscanf(line, "%s %d", &direction, &steps)
		for i := 0; i < steps; i++ {

			switch direction {
			case "U":
				snake[0].x++
			case "D":
				snake[0].x--
			case "R":
				snake[0].y++
			case "L":
				snake[0].y--
			}

			for j := 1; j < N; j++ {
				if IsTooFar(snake[j-1], snake[j]) {
					snake[j] = MoveTo(snake[j-1], snake[j])
				}
			}
			visited[snake[N-1]] = true
		}
	}
	fmt.Println(len(visited))
}

func IsTooFar(head Index, tail Index) bool {

	dx := math.Abs(float64(head.x - tail.x))
	dy := math.Abs(float64(tail.y - head.y))

	return dx > 1 || dy > 1
}

func MoveTo(head Index, tail Index) Index {
	dx := math.Abs(float64(head.x - tail.x))
	dy := math.Abs(float64(tail.y - head.y))

	dirx := 0
	diry := 0
	if dx > 0 {
		dirx = (head.x - tail.x) / int(dx)
	}
	if dy > 0 {
		diry = (head.y - tail.y) / int(dy)
	}

	return Index{tail.x + dirx, tail.y + diry}
}
