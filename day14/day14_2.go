package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("./day10/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	register := 1
	cycle := 1

	crt_widht := 40
	crt_height := 6
	var crt = make([]string, 6)

	amount := 0
	read := true
	var line string

	for cycle <= crt_widht*crt_height {

		i := (cycle - 1) / crt_widht
		j := (cycle - 1) % crt_widht

		if j <= register-1+2 && j >= register-1 {
			crt[i] += "#"
		} else {
			crt[i] += "."
		}

		if read {
			sc.Scan()
			line = sc.Text()
			line = strings.TrimSpace(line)

			data := strings.Split(line, " ")

			if len(data) == 2 {
				amount, _ = strconv.Atoi(data[1])
				read = false
			}
			cycle++
		} else {
			register += amount
			read = true
			cycle++
		}
	}

	for i := 0; i < crt_height; i++ {
		fmt.Println(crt[i])
	}
}
