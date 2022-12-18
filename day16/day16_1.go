package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Note struct {
	earn_value int
	opened     string
}

func main() {
	INPUT_PATH := "./day16/input.txt"
	graph, flows := read_input_day16(INPUT_PATH)

	//to_remove := map[string]bool{}
	//
	//for name, neighbors := range graph {
	//	if name == "AA" {
	//		continue
	//	}
	//
	//	if flows[name] == 0 {
	//		for n1, d1 := range neighbors {
	//			for n2, d2 := range neighbors {
	//				if n1 == n2 {
	//					continue
	//				}
	//				graph[n1][n2] = d1 + d2
	//				graph[n2][n1] = d1 + d2
	//			}
	//		}
	//		to_remove[name] = true
	//	}
	//}
	//
	//for name, _ := range to_remove {
	//	delete(graph, name)
	//}
	//for name, _ := range graph {
	//	for n, _ := range graph[name] {
	//		if to_remove[n] {
	//			delete(graph[name], n)
	//		}
	//	}
	//}

	//fmt.Println(graph)
	max_flow := -1
	total_flow := 0
	for _, v := range flows {
		if max_flow < v {
			max_flow = v
		}
		total_flow += v
	}

	// map [time][valve][opened]  value
	table := map[int]map[string]map[string]int{}
	for i := 0; i < 30; i++ {
		table[i+1] = map[string]map[string]int{}
	}

	table[1]["AA"] = map[string]int{}
	table[1]["AA"][""] = 0

	for t := 2; t <= 30; t++ {

		for name, neighbors := range graph {
			flow := flows[name]
			table[t][name] = map[string]int{}
			//m := table[t-1][name] // do nothing stay where was

			for path, value := range table[t-1][name] {
				table[t][name][path] = value
			}
			// open valve
			for path, value := range table[t-1][name] {
				if !strings.Contains(path, name) && name != "AA" && flow > 0 {
					value = value + flow*(30-t+1)
					if path != "" {
						path += "-"
					}
					//fmt.Println(path)
					path += name
					valves := strings.Split(path, "-")
					sort.Strings(valves)
					//fmt.Println(valves)
					path = strings.Join(valves, "-")
					//fmt.Println(path)

					table[t][name][path] = intMax(table[t][name][path], value)
				}
			}

			// maximum of neighbors
			for neighbor, d := range neighbors {
				if t-d < 1 {
					continue
				}
				n3 := table[t-d][neighbor]
				for path, value := range n3 {
					table[t][name][path] = intMax(table[t][name][path], value)
				}
			}
		}
	}
	//fmt.Println(table[30])

	max := 0

	for _, n := range table[30] {
		for _, v := range n {
			if v > max {
				max = v
			}
		}
	}
	fmt.Println(max)
	fmt.Println(max_flow)
}

func read_input_day16(INPUT_PATH string) (map[string]map[string]int, map[string]int) {
	input, _ := os.Open(INPUT_PATH)
	defer input.Close()
	sc := bufio.NewScanner(input)

	graph := map[string]map[string]int{}
	flows := map[string]int{}

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)
		data := strings.Split(line, " ")

		name := data[1]

		var flow int
		fmt.Sscanf(data[4], "rate=%d;", &flow)
		flows[name] = flow
		// flow rate at 4
		// to 9+
		neighbors := map[string]int{}
		for i := 9; i < len(data); i++ {
			var n string
			if i+1 == len(data) {
				n = data[i]
			} else {
				n = data[i][:len(data[i])-1]
			}

			neighbors[n] = 1

		}

		graph[name] = neighbors

	}
	return graph, flows
}

func NoteMax(a, b Note) Note {
	if a.earn_value < b.earn_value {
		return b
	}
	if a.earn_value == b.earn_value && a.opened != b.opened {
		fmt.Println("++++++", a.opened, b.opened)
	}
	return a
}

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
