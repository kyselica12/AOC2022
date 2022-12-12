package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Monkey struct {
	items            []int
	increment        func(int) int
	divisible        int
	monkey1, monkey2 int
	inspected        int
}

func main() {
	input, _ := os.Open("./day11/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	nRounds := 10000
	var monkeys []Monkey

	for {

		data, ok := get_line(sc)
		if !ok {
			break
		}
		if len(data) > 0 && data[0] == "Monkey" {
			monkey := readMonkey(data, sc)
			monkeys = append(monkeys, monkey)
		}

	}

	modulo := 1

	for _, m := range monkeys {
		modulo *= m.divisible
	}

	fmt.Println(monkeys)
	//max := 0
	for i := 0; i < nRounds; i++ {
		for j, m := range monkeys {
			for _, item := range m.items {
				value := m.increment(item)
				value = value % modulo
				//if value > max {
				//	max = value
				//	fmt.Println(value)
				//}
				monkeys[j].inspected++
				//fmt.Print(" value ", value, item)
				if value%m.divisible == 0 {
					monkeys[m.monkey1].items = append(monkeys[m.monkey1].items, value)
				} else {
					monkeys[m.monkey2].items = append(monkeys[m.monkey2].items, value)
				}
			}
			monkeys[j].items = []int{}
			//fmt.Println(monkeys[j].items)
		}
		//fmt.Println(i, monkeys)

		//break
	}

	var vals []int
	for _, m := range monkeys {
		vals = append(vals, m.inspected)
	}

	sort.Ints(vals)
	fmt.Println(vals)
	fmt.Println(vals[len(vals)-1] * vals[len(vals)-2])
}

func readMonkey(data []string, sc *bufio.Scanner) Monkey {
	data, _ = get_line(sc)
	var items []int
	for i, item := range data[2:] {
		if i < len(data[2:])-1 {
			item = item[:len(item)-1]
		}
		atoi, _ := strconv.Atoi(item[:len(item)])
		items = append(items, atoi)
	}

	data, _ = get_line(sc)

	sign := data[4]

	value, err := strconv.Atoi(data[5])

	var increment func(int) int
	switch sign {
	case "*":
		increment = func(i int) int {
			if err == nil {
				return i * value
			}
			return i * i
		}
	case "+":
		increment = func(i int) int {
			if err == nil {
				return i + value
			}
			return i + i
		}
	}

	data, _ = get_line(sc)
	condition_value, _ := strconv.Atoi(data[3])

	data, _ = get_line(sc)
	m1, _ := strconv.Atoi(data[5])
	data, _ = get_line(sc)
	m2, _ := strconv.Atoi(data[5])

	monkey := Monkey{items: items, increment: increment, divisible: condition_value, monkey1: m1, monkey2: m2}
	return monkey
}

func get_line(sc *bufio.Scanner) ([]string, bool) {
	ok := sc.Scan()
	if ok {
		line := sc.Text()
		line = strings.TrimSpace(line)
		data := strings.Split(line, " ")
		return data, true
	}
	return nil, false
}
