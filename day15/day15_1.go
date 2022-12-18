package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func main() {
	input, _ := os.Open("./day15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	scans := map[Pos]int{}
	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)

		var xs, ys, xb, yb int
		n, _ := fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		fmt.Println(n, line)
		scans[Pos{xs, ys}] = int(math.Abs(float64(xs-xb)) + math.Abs(float64(ys-yb)))

	}

	y := 2000000

	var start, end int
	start = math.MaxInt
	end = math.MinInt
	for s, dist := range scans {
		dy := int(math.Abs(float64(s.y - y)))
		//fmt.Println(s, dist, dy)
		dist -= dy
		if dist < 0 {
			continue
		}
		if start > s.x-dist {
			start = s.x - dist
		}
		if end < s.x+dist {
			end = s.x + dist
		}
	}

	fmt.Println(start, end, end-start)

}
