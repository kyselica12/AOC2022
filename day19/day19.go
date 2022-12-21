package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type State struct {
	resources, robots []int
	t                 int
}

func main() {
	PATH := "./day19/input.txt"
	blue_prints := read_blueprints(PATH)[:3]

	fmt.Println(blue_prints)
	result := 1

	c_results := make(chan int, len(blue_prints))

	for id := 0; id < len(blue_prints); id++ {
		blueprint := blue_prints[id]

		go search(blueprint, id, c_results)
	}
	for i := 0; i < len(blue_prints); i++ {
		x := <-c_results
		fmt.Println("X: ", x)
		result *= x
	}
	fmt.Println("Result: ", result)
	//fmt.Println(search(24, resources, robots, blue_prints[0], map[string]int{}))
}

func search(blueprint [][]int, id int, ch chan<- int) {
	queue := []State{{[]int{0, 0, 0, 0}, []int{1, 0, 0, 0}, 32}}
	depth := 32

	visited := map[string]bool{}
	max := 0

	max_curr := 0
	max_prev := 0

	var max_cost []int
	for i := 0; i < 3; i++ {
		m := 0
		for j := 0; j < 4; j++ {
			m = maxInt(m, blueprint[j][i])
		}
		max_cost = append(max_cost, m)
	}
	max_cost = append(max_cost, math.MaxInt)

	for len(queue) > 0 {

		s := queue[0]
		queue = queue[1:]
		//fmt.Println(s)
		if depth > s.t {
			depth--
			//max_prev = max_curr
			fmt.Print(depth, " ")
		}

		if s.t == 0 {
			max = maxInt(max, s.resources[3])
			continue
		}

		x := s.resources[3] + s.t*s.robots[3] + ((s.t+1)/2)*s.t
		max_curr = maxInt(max_curr, x)

		if x < max_prev {
			continue
		}

		//cut down unwantable resources
		for i := 0; i < 3; i++ {
			m := max_cost[i]
			if s.resources[i] >= m*s.t-(s.t-1)*s.robots[i] {
				s.resources[i] = m*s.t - (s.t-1)*s.robots[i]
			}
		}

		afford_all := true
		for robot_id, recipe := range blueprint {
			can_affort := true

			if s.robots[robot_id] >= max_cost[robot_id] {
				can_affort = false
			}

			for j := 0; j < 4; j++ {
				if can_affort {
					if s.resources[j] < recipe[j] {
						can_affort = false
						afford_all = false
					}
				}
			}

			if can_affort {
				var resources2, robots2 []int

				resources2 = append(resources2, s.resources...)
				robots2 = append(robots2, s.robots...)

				for i := 0; i < 4; i++ {
					resources2[i] -= blueprint[robot_id][i]
					resources2[i] += s.robots[i]
				}

				robots2[robot_id]++

				new_state := State{resources2, robots2, s.t - 1}

				key := getKey(new_state)
				if visited[key] {
					continue
				}
				visited[key] = true
				queue = append(queue, new_state)
			}
		}
		if !afford_all {
			var resources2, robots2 []int
			resources2 = append(resources2, s.resources...)
			robots2 = append(robots2, s.robots...)

			for i := 0; i < 4; i++ {
				resources2[i] += s.robots[i]
			}

			new_state := State{resources2, robots2, s.t - 1}
			visited[getKey(new_state)] = true
			queue = append(queue, new_state)
		}
	}
	fmt.Println("max: ", max, id)
	ch <- max
}

func getKey(s State) string {
	return fmt.Sprintf("%q", s.resources) + fmt.Sprintf("%q", s.robots) + strconv.Itoa(s.t)
}

func maxInt(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func read_blueprints(PATH string) [][][]int {
	input, _ := os.Open(PATH)
	defer input.Close()
	sc := bufio.NewScanner(input)

	var blue_prints [][][]int

	for sc.Scan() {

		line := sc.Text()
		line = strings.TrimSpace(line)

		data := strings.Split(line, ":")
		sentences := strings.Split(data[1], ".")

		for i, s := range sentences {
			sentences[i] = strings.TrimSpace(s)
		}
		bp := make([][]int, 4)
		for i := 0; i < 4; i++ {
			bp[i] = make([]int, 4)
		}
		var nOre, nClay, nObsidian int

		fmt.Sscanf(sentences[0], "Each ore robot costs %d ore", &nOre)
		bp[0][0] = nOre

		fmt.Sscanf(sentences[1], "Each clay robot costs %d ore", &nOre)
		bp[1][0] = nOre

		fmt.Sscanf(sentences[2], "Each obsidian robot costs %d ore and %d clay", &nOre, &nClay)
		bp[2][0] = nOre
		bp[2][1] = nClay

		fmt.Sscanf(sentences[3], "Each geode robot costs %d ore and %d obsidian ", &nOre, &nObsidian)
		bp[3][0] = nOre
		bp[3][2] = nObsidian

		blue_prints = append(blue_prints, bp)

	}
	return blue_prints
}
