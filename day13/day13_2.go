package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("./day13/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var packets []any

	packets = read_file(sc)
	deliminator1 := readLine("[[2]]")
	deliminator2 := readLine("[[6]]")
	fmt.Println(deliminator2, deliminator1)

	sort.Slice(packets, func(i, j int) bool {
		return compare(packets[i], packets[j]) == 1
	})
	x := -1
	y := -1

	for i, s := range packets {
		if compare(s, deliminator1) == 0 {
			x = i + 1
			fmt.Println("x", s, i+1)
		}
		if compare(s, deliminator2) == 0 {
			y = i + 1
			fmt.Println("y", s, i+1)
		}
	}

	fmt.Println(x, y, x*y)
}

func compare(b1 interface{}, b2 interface{}) int {
	//fmt.Println(b1)
	//fmt.Println(compare(b1,b2))

	if v1, ok := b1.(int); ok {
		if v2, ok2 := b2.(int); ok2 {
			if v1 < v2 {
				return 1
			} else if v1 == v2 {
				return 0
			}
			return -1
		}
		b1 = []interface{}{b1}
	}
	if _, ok := b2.(int); ok {
		b2 = []interface{}{b2}
	}
	arr1 := b1.([]interface{})
	arr2 := b2.([]interface{})
	for i := 0; i < len(arr1); i++ {
		if i >= len(arr2) {
			return -1
		}
		c := compare(arr1[i], arr2[i])
		if c != 0 {
			return c
		}
	}
	if len(arr1) == len(arr2) {
		return 0
	}
	return 1
}

func readLine(line1 string) interface{} {
	line := []rune(line1)
	_, block := read_one(line, 0)
	return block
}

func read_one(line []rune, i int) (int, interface{}) {

	if line[i] == ',' {
		return read_one(line, i+1)
	}

	if line[i] == '[' {
		return read_array(line, i)
	}

	return read_int(line, i)
}

func read_int(line []rune, i int) (int, interface{}) {
	number := ""
	j := i
	for line[j] >= '0' && line[j] <= '9' {
		number += string(line[j])
		j++
	}

	b, _ := strconv.Atoi(number)

	if j < len(line) && line[j] == ',' {
		return j + 1, b
	}

	return j, b
}

func read_array(line []rune, i int) (int, interface{}) {
	j := i + 1

	var arr []interface{}
	for j < len(line) && line[j] != ']' {

		k, b := read_one(line, j)
		arr = append(arr, b)
		j = k
	}
	return j + 1, arr
}

func read_file(sc *bufio.Scanner) []any {
	var line string
	var packets []any
	for sc.Scan() {
		line = sc.Text()
		if len(line) > 0 {
			line = strings.TrimSpace(line)
			packets = append(packets, readLine(line))
		}
	}
	return packets
}
