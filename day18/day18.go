package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
	x, y, z int
}

func main() {
	input, _ := os.Open("./day18/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	cubes := map[Pos]bool{}
	surface := 0

	var xs []int
	var ys []int
	var zs []int

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)
		var x, y, z int
		fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)

		cubes[Pos{x, y, z}] = true

		xs = append(xs, x)
		ys = append(ys, y)
		zs = append(zs, z)

		neighbors := get_neighbors(x, y, z)
		for _, n := range neighbors {
			if cubes[n] {
				surface--
			} else {
				surface++
			}
		}

	}

	fmt.Println("Pt1: ", surface)

	min_x, max_x := get_bounds(xs)
	min_y, max_y := get_bounds(ys)
	min_z, max_z := get_bounds(zs)

	for x := min_x; x < max_x+1; x++ {
		for y := min_y; y < max_y+1; y++ {
			for z := min_z; z < max_z+1; z++ {
				id := Pos{x, y, z}
				if !cubes[id] && !isWayOut(x, y, z, min_x, max_x, min_y, max_y, min_z, max_z, cubes) {
					for _, n := range get_neighbors(x, y, z) {
						if cubes[n] {
							surface--
						}
					}
				}
			}
		}
	}

	fmt.Println("Pt2: ", surface)

}

func isWayOut(x, y, z, min_x, max_x, min_y, max_y, min_z, max_z int, cubes map[Pos]bool) bool {

	queue := []Pos{{x, y, z}}
	visited := map[Pos]bool{}
	visited[Pos{x, y, z}] = true

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]

		if p.x < min_x || max_x < p.x || p.y < min_y || max_y < p.y || p.z < min_z || max_z < p.z {
			//fmt.Println("Out -> ", Pos{x, y, z}, p)
			return true
		}

		for _, n := range get_neighbors(p.x, p.y, p.z) {
			if !visited[n] && !cubes[n] {
				queue = append(queue, n)
				visited[n] = true
			}
		}
	}
	return false
}

func get_neighbors(x int, y int, z int) []Pos {
	neighbors := []Pos{
		Pos{x - 1, y, z},
		Pos{x + 1, y, z},
		Pos{x, y - 1, z},
		Pos{x, y + 1, z},
		Pos{x, y, z - 1},
		Pos{x, y, z + 1},
	}
	return neighbors
}

func get_bounds(s []int) (int, int) {

	min := math.MaxInt
	max := math.MinInt

	for _, x := range s {
		if min > x {
			min = x
		}
		if max < x {
			max = x
		}
	}

	return min, max
}
