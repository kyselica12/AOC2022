package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	input, _ := os.Open("./day8/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)
	visible := 0
	var forest [][]rune

	for sc.Scan() {
		line := sc.Text()
		line = strings.TrimSpace(line)
		line_chars := []rune(line)
		visible += len(line_chars)
		forest = append(forest, line_chars)
	}

	not_visible := map[int][]int{}

	for r, line := range forest[1 : len(forest)-1] {
		max_tree := '0' - 1
		for c, tree := range line {
			if max_tree >= tree {
				not_visible[r+1] = append(not_visible[r+1], c)
			}
			if max_tree < tree {
				max_tree = tree
			}
		}
	}
	//fmt.Println(not_visible)
	not_visible_column := map[int][]int{}
	for r, cols := range not_visible {

		max := '0' - 1
		prev := len(forest[r])
		for i := len(cols) - 1; i >= 0; i-- {
			c := cols[i]
			line := forest[r]
			max = MaxSlice(line[c+1:prev], max)
			if max >= line[c] {
				not_visible_column[c] = append(not_visible_column[c], r)
			} else {
				max = line[c]
			}
		}
	}

	//fmt.Println(not_visible_column)

	for c, rows := range not_visible_column {
		max := '0' - 1
		prev := 0
		//fmt.Println(rows)
		sort.Ints(rows)
		var rows_cadidates = []int{}
		for _, r := range rows {
			max = MaxColumnDown(c, prev, r, max, forest)
			if max >= forest[r][c] {
				rows_cadidates = append(rows_cadidates, r)
			} else {
				max = forest[r][c]
			}
		}

		max = '0' - 1
		prev = len(forest[c])
		//fmt.Println(rows_cadidates)
		for i := len(rows_cadidates) - 1; i >= 0; i-- {
			r := rows_cadidates[i]
			//println(r, prev)
			max = MaxColumnDown(c, r+1, prev, max, forest)
			if max >= forest[r][c] {
				//fmt.Println(r, c, string(forest[r][c]))
				visible--
			} else {
				max = forest[r][c]
			}
		}
		//println()

	}

	//fmt.Println(not_visible)
	//fmt.Println(not_visible_column)

	fmt.Println(visible)
}

func MaxColumnDown(index int, start int, end int, max rune, forest [][]rune) rune {
	for i := start; i < end; i++ {
		//fmt.Println(len(forest), len(forest[index]))
		tree := forest[i][index]
		if tree > max {
			max = tree
		}
	}
	return max
}

func MaxSlice(v []rune, max rune) rune {
	for _, x := range v {
		if max < x {
			max = x
		}
	}
	return max
}
