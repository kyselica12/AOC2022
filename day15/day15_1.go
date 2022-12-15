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
	input, _ := os.Open("./day15/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)

		fmt.Sscanf()

	}

	fmt.Println(0)

}
