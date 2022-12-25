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
	PATH := "./day23/input.txt"
	elfs := read_file_day23(PATH)

	//fmt.Println("Elfs ", len(elfs))
	move_order := []rune{'N', 'S', 'W', 'E'}
	round := 0
	for {
		round++
		//fmt.Println("================ ROUND ", round, " ===================")
		//print_board(elfs)
		move_count := map[Pos]int{}
		moves := map[Pos]Pos{}

		//fmt.Println(string(move_order))

		for p, ok := range elfs {

			if !ok {
				continue
			}
			//pp := Pos{1, 3}
			//if p == pp {
			//	fmt.Println("SAAAAA")
			//}

			move := p

			nort := Pos{p.x - 1, p.y}
			nort_east := Pos{p.x - 1, p.y + 1}
			nort_west := Pos{p.x - 1, p.y - 1}
			west := Pos{p.x, p.y - 1}
			east := Pos{p.x, p.y + 1}
			south := Pos{p.x + 1, p.y}
			south_east := Pos{p.x + 1, p.y + 1}
			south_west := Pos{p.x + 1, p.y - 1}

			if elfs[nort] || elfs[nort_east] || elfs[nort_west] || elfs[west] || elfs[south_west] || elfs[south] || elfs[south_east] || elfs[east] {
				move, ok = find_move(elfs, nort, nort_west, nort_east, south, south_east, south_west, west, east, move_order)
				if !ok {
					move = p
				}
			}
			move_count[move]++
			if move != p {
				moves[p] = move
			}
		}
		if len(moves) == 0 {
			break
		}

		for from, to := range moves {
			if move_count[to] > 1 || from == to {
				continue
			}
			delete(elfs, from)
			elfs[to] = true
		}

		if round%100 == 0 {
			fmt.Println("Round: ", round, len(moves))
		}

		move_order = append(move_order, move_order[0])[1:]
	}
	//fmt.Println("================ ROUND ", 10, " ===================")
	//print_board(elfs)
	//minx, maxx, miny, maxy, count := get_bounds(elfs)
	//fmt.Println(count)
	//fmt.Println((maxx-minx+1)*(maxy-miny+1) - count)
	fmt.Println(round)

}

func find_move(elfs map[Pos]bool, nort Pos, nort_west Pos, nort_east Pos,
	south Pos, south_east Pos, south_west Pos, west Pos, east Pos,
	move_order []rune) (Pos, bool) {

	var move Pos
	var ok bool
	for _, c := range move_order {
		switch c {
		case 'N':
			move, ok = move_to(nort, nort_west, nort_east, elfs)
		case 'S':
			move, ok = move_to(south, south_west, south_east, elfs)
		case 'W':
			move, ok = move_to(west, south_west, nort_west, elfs)
		case 'E':
			move, ok = move_to(east, nort_east, south_east, elfs)
		}
		if ok {
			return move, true
		}
	}
	return move, false
}

func move_to(dest_pos, pos2, pos3 Pos, elfs map[Pos]bool) (Pos, bool) {
	if !elfs[dest_pos] && !elfs[pos2] && !elfs[pos3] {
		return dest_pos, true
	}
	return dest_pos, false
}

func get_bounds(elfs map[Pos]bool) (int, int, int, int, int) {
	minx := math.MaxInt
	maxx := math.MinInt
	miny := math.MaxInt
	maxy := math.MinInt
	count := 0
	for p, ok := range elfs {
		count++
		if !ok {
			continue
		}
		if p.x < minx {
			minx = p.x
		}
		if p.x > maxx {
			maxx = p.x
		}
		if p.y < miny {
			miny = p.y
		}
		if p.y > maxy {
			maxy = p.y
		}
	}
	return minx, maxx, miny, maxy, count
}

func print_board(elfs map[Pos]bool) {

	minx, maxx, miny, maxy, _ := get_bounds(elfs)

	for i := minx; i <= maxx; i++ {
		for j := miny; j <= maxy; j++ {
			if elfs[Pos{i, j}] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func read_file_day23(PATH string) map[Pos]bool {
	input, _ := os.Open(PATH)
	defer input.Close()
	sc := bufio.NewScanner(input)

	elfs := map[Pos]bool{}
	i := 0
	for sc.Scan() {
		line := sc.Text()
		striped_line := strings.TrimSpace(line)

		for j := 0; j < len(striped_line); j++ {
			if striped_line[j] == '#' {
				elfs[Pos{i, j}] = true
			}
		}

		i++
	}
	fmt.Println(len(elfs))
	return elfs
}
