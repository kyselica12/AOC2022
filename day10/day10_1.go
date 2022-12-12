package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("./day10/test-input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	register := 1
	cycle := 1
	total := 0
	//nline := 0
	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)
		//nline++
		data := strings.Split(line, " ")

		if len(data) == 1 {
			cycle++
			if (cycle-20)%40 == 0 {
				total += cycle * register
				fmt.Println(cycle, register, line)
			}
		} else {
			amount, _ := strconv.Atoi(data[1])
			cycle += 2
			register += amount
			if (cycle-20)%40 == 0 {
				total += cycle * register
				fmt.Println(cycle, register, line)
			} else if (cycle-20)%40 == 1 {
				fmt.Println(cycle-1, register-amount, line)
				total += (cycle - 1) * (register - amount)
			}
		}

	}
	fmt.Println(total)
}
