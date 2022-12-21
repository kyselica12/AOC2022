package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type number int

type operation struct {
	m1, m2  string
	operant string
}

func main() {
	input, _ := os.Open("./day21/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	monkeys := map[string]any{}

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)

		data := strings.Split(line, " ")
		name := data[0][:len(data[0])-1]
		if len(data) == 2 {
			x, _ := strconv.Atoi(data[1])
			monkeys[name] = x
		} else {
			monkeys[name] = operation{data[1], data[3], data[2]}
		}
	}

	//fmt.Println("Part1:", compute("root", monkeys))
	root := monkeys["root"].(operation)
	right := compute(root.m2, monkeys)

	res := compute_eq(root.m1, monkeys, right)

	fmt.Println(res)

}

func compute(monkey string, monkeys map[string]any) int {

	if monkey == "humn" {
		fmt.Println("HUUUUUUUUUUMN")
	}

	if v, ok := monkeys[monkey].(int); ok {
		return v
	}
	op := monkeys[monkey].(operation)
	n1 := compute(op.m1, monkeys)
	n2 := compute(op.m2, monkeys)
	switch op.operant {
	case "+":
		return n1 + n2
	case "-":
		return n1 - n2
	case "*":
		return n1 * n2
	case "/":
		return n1 / n2
	}
	return -1
}

func compute_eq(name string, monkeys map[string]any, right int) int {

	if name == "humn" {
		return right
	}

	if _, ok := monkeys[name].(int); ok {
		return 0
	}

	op := monkeys[name].(operation)

	var value int
	found := op.m1
	if contains_humn(op.m1, monkeys) {
		value = compute(op.m2, monkeys)
	} else {
		found = op.m2
		value = compute(op.m1, monkeys)
	}

	switch op.operant {
	case "+":
		return compute_eq(found, monkeys, right-value)
	case "-":
		if found == op.m1 {
			return compute_eq(found, monkeys, right+value)
		}
		return compute_eq(found, monkeys, value-right)
	case "*":
		return compute_eq(found, monkeys, right/value)
	case "/":
		if found == op.m1 {
			return compute_eq(found, monkeys, right*value)
		}
		return compute_eq(found, monkeys, value/right)
	}
	return -1
}

func contains_humn(name string, monkeys map[string]any) bool {

	if name == "humn" {
		return true
	}

	if _, ok := monkeys[name].(int); ok {
		return false
	}
	op := monkeys[name].(operation)

	return contains_humn(op.m1, monkeys) || contains_humn(op.m2, monkeys)
}
