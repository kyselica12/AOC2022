package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	PATH := "./day25/input.txt"
	input, _ := os.Open(PATH)
	defer input.Close()

	sc := bufio.NewScanner(input)
	sum := 0
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		value := snafu_to_int(line)
		//fmt.Println(line, value)
		sum += value
	}
	n := sum
	snafu := int_to_snafu(n)
	//fmt.Println(n)
	//fmt.Println(snafu)
	//fmt.Println(snafu_to_int(snafu))

	fmt.Println("Value: ", sum, " snafu: ", snafu)

}

func int_to_snafu(n int) string {
	snafu := ""
	over := 0
	for n > 0 {
		//fmt.Println("V: ", n, " mod: ", n%5, " over: ", over)
		v := (n + over) % 5
		if v > 2 || (v == 0 && over == 1) {
			over = 1
		} else {
			over = 0
		}
		switch v {
		case 0:
			snafu = "0" + snafu
		case 1:
			snafu = "1" + snafu
		case 2:
			snafu = "2" + snafu
		case 3:
			snafu = "=" + snafu
		case 4:
			snafu = "-" + snafu
		}
		n = (n - n%5) / 5
	}
	if over == 1 {
		snafu = "1" + snafu
	}
	return snafu
}

func snafu_to_int(line string) int {
	value := 0
	base := 1
	for i := len(line) - 1; i >= 0; i-- {
		switch line[i] {
		case '2':
			value += base * 2
		case '1':
			value += base
		case '-':
			value -= base
		case '=':
			value -= 2 * base
		}
		base *= 5
	}
	return value
}
