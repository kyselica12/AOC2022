package main

import (
	"bufio"
	"fmt"
	"github.com/albertorestifo/dijkstra"
	"os"
	"sort"
	"strings"
)

func main() {
	INPUT_PATH := "./day16/input.txt"
	graph, flows := read_input_day16(INPUT_PATH)

	to_remove := map[string]bool{}

	for name, neighbors := range graph {
		if name == "AA" {
			continue
		}

		if flows[name] == 0 {
			for n1, d1 := range neighbors {
				for n2, d2 := range neighbors {
					if n1 == n2 {
						continue
					}
					graph[n1][n2] = d1 + d2
					graph[n2][n1] = d1 + d2
				}
			}
			to_remove[name] = true
		}
	}

	for name, _ := range to_remove {
		delete(graph, name)
	}
	for name, _ := range graph {
		for n, _ := range graph[name] {
			if to_remove[n] {
				delete(graph[name], n)
			}
		}
	}

	dist := map[string]map[string]int{}
	G := dijkstra.Graph(graph)

	for n1, _ := range graph {
		dist[n1] = map[string]int{}
		for n2, _ := range graph {
			_, d, _ := G.Path(n1, n2)
			dist[n1][n2] = d
		}
	}
	//fmt.Println(dist)

	memory := map[string]int{}
	res := search(26, "AA", "", dist, flows, memory, true)
	//fmt.Println(memory)
	fmt.Println(res)
}

func search(t int, pos string, visited string, dist map[string]map[string]int, flows map[string]int,
	memory map[string]int, elephant bool) int {
	max := 0
	id := string(t) + pos + visited
	if elephant {
		id += "E"
	}
	if v, ok := memory[id]; ok {
		return v
	}
	for n1, _ := range dist {
		if strings.Contains(visited, n1) {
			continue
		}

		d := dist[pos][n1]
		if d > t {
			continue
		}
		//fmt.Println(t, d, n1, pos)
		new_visited := visited
		if visited != "" {
			new_visited += "-"
		}
		new_visited += n1
		valves := strings.Split(new_visited, "-")
		sort.Strings(valves)
		new_visited = strings.Join(valves, "-")

		val := flows[n1]*(t-d-1) + search(t-d-1, n1, new_visited, dist, flows, memory, elephant)
		max = intMax(max, val)
		if elephant {
			elephant_val := search(26, "AA", visited, dist, flows, memory, false)
			max = intMax(max, elephant_val)
		}
	}

	memory[id] = max
	return max
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

func intMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
