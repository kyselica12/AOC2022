package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input, _ := os.Open("./day6/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	sc.Split(bufio.ScanRunes)
	buff := make([]string, 14)

	index := 0
	for sc.Scan() {

		x := sc.Text()
		if index < 14 {
			buff[index] = x
		} else {
			buff = append(buff[1:], x)
		}

		fmt.Println(buff)
		index++

		if index > 13 && check(buff) {
			fmt.Println(index)
			break
		}
	}

}

func add(buff []rune, elem rune, index int) []rune {
	if index < 14 {
		buff[index] = elem
		return buff
	}

	buff = buff[1:]
	buff = append(buff, elem)
	return buff
}

func check(buff []string) bool {
	unique := map[string]bool{}
	for i := 0; i < 14; i++ {
		if unique[buff[i]] {
			return false
		}
		unique[buff[i]] = true
	}
	return true
}
