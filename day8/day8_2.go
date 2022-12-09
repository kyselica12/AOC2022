package main

import (
	"bufio"
	"fmt"
	"os"
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

	max := -1

	for r := 0; r < len(forest); r++ {
		for c := 0; c < len(forest[r]); c++ {
			val := getScore(r, c, forest)
			//println(r, c, val)
			if val > max {
				max = val
			}
		}
	}
	//getScore(3, 2, forest)
	fmt.Println(max)

}

func getScore(r int, c int, forest [][]rune) int {

	tree := forest[r][c]
	var down, up, left, right int

	for i := r + 1; i < len(forest); i++ {
		down++
		//fmt.Println("Down", forest[i][c], i, c)
		if forest[i][c] >= tree {
			break
		}
	}

	for i := r - 1; i > -1; i-- {
		up++
		//fmt.Println("Up", forest[i][c], i, c)
		if forest[i][c] >= tree {
			break
		}
	}

	for i := c + 1; i < len(forest[r]); i++ {
		left++
		//fmt.Println("Left", forest[r][i], r, i)
		if forest[r][i] >= tree {
			break
		}
	}

	for i := c - 1; i > -1; i-- {
		right++
		//fmt.Println("Right", forest[r][i], r, i)
		if forest[r][i] >= tree {
			break
		}
	}

	//fmt.Println(r, c, down, up, left, right)

	return down * up * left * right
}
