package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

type Interval struct {
	borders []int
}

func (interval *Interval) add(start, end int) {
	if len(interval.borders) == 0 {
		interval.borders = append(interval.borders, start)
		interval.borders = append(interval.borders, end)
	} else {

		extend := -1
		insert := -1
		for i := 0; i < len(interval.borders); i += 2 {
			istart := interval.borders[i]
			iend := interval.borders[i+1]

			if (istart <= end && end <= iend) || (istart <= start && start <= iend) || (start < istart && end > iend) {
				extend = i
				break
			}

			if end < istart {
				insert = i
				break
			}
		}

		if insert != -1 {
			var new_borders []int
			new_borders = append(new_borders, interval.borders[:insert]...)
			new_borders = append(new_borders, []int{start, end}...)
			new_borders = append(new_borders, interval.borders[insert:]...)
			interval.borders = new_borders
		} else if extend != -1 {
			istart := interval.borders[extend]
			iend := interval.borders[extend+1]
			interval.borders[extend] = int(math.Min(float64(istart), float64(start)))
			interval.borders[extend+1] = int(math.Max(float64(iend), float64(end)))

			var new_borders []int

			prev_start := interval.borders[0]
			prev_end := interval.borders[1]
			for i := 2; i < len(interval.borders); i += 2 {
				if prev_end < interval.borders[i] {
					new_borders = append(new_borders, []int{prev_start, prev_end}...)
					prev_start = interval.borders[i]
					prev_end = interval.borders[i+1]
				} else {
					prev_end = int(math.Max(float64(interval.borders[i+1]), float64(prev_end)))
				}
			}
			new_borders = append(new_borders, []int{prev_start, prev_end}...)

			interval.borders = new_borders
		} else {
			interval.borders = append(interval.borders, []int{start, end}...)
		}

	}
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
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &xs, &ys, &xb, &yb)
		scans[Pos{xs, ys}] = int(math.Abs(float64(xs-xb)) + math.Abs(float64(ys-yb)))

	}

	t := time.Now()

	N := 4000000
	for y := 0; y <= N; y++ {

		interval := Interval{}

		for s, dist := range scans {
			dy := int(math.Abs(float64(s.y - y)))
			if dist-dy < 0 {
				continue
			}
			dist -= dy

			start := int(math.Max(0, float64(s.x-dist)))
			end := int(math.Min(float64(N), float64(s.x+dist)))

			interval.add(start, end)

		}
		if len(interval.borders) > 2 {
			fmt.Println(y, len(interval.borders), interval.borders, "\n")
			x := interval.borders[1] + 1
			fmt.Println("res: ", x*N+y)
			fmt.Println("time: ", time.Now().Sub(t))
			break
		}
	}
}
