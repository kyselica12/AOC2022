package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	value int
	next  *Node
	prev  *Node
}

type Item struct {
	value, index int
}

func main() {

	PATH := "./day20/input.txt"
	file, _ := os.Open(PATH)

	sc := bufio.NewScanner(file)
	var data []int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		n, _ := strconv.Atoi(line)
		data = append(data, n*811589153)
	}

	var items []Item

	for i, n := range data {
		items = append(items, Item{n, i})
	}

	slice_list(data, 10)
	fmt.Println("------------------")
	linked_list(data)

}

func slice_list(data []int, nrot int) {
	var items []Item
	for i, n := range data {
		items = append(items, Item{n, i})
	}

	rotated_items := append([]Item{}, items...)
	N := len(items)
	for rot := 0; rot < nrot; rot++ {

		for _, item := range items {
			if item.value == 0 {
				continue
			}
			index := 0 // rottated_items.index(item)
			for i := 0; i < len(items); i++ {
				if rotated_items[i] == item {
					break
				}
				index++
			}

			dst := (N - 1 + (index+item.value)%(N-1)) % (N - 1)

			tmp := append([]Item{}, rotated_items[:index]...)
			tmp = append(tmp, rotated_items[index+1:]...)

			rotated_items = append([]Item{}, tmp[:dst]...)
			rotated_items = append(rotated_items, item)
			rotated_items = append(rotated_items, tmp[dst:]...)
		}
	}
	index := 0 // rottated_items.index(item)
	for i := 0; i < len(items); i++ {
		if rotated_items[i].value == 0 {
			break
		}
		index++
	}

	v1 := rotated_items[(index+1000)%N].value
	v2 := rotated_items[(index+2000)%N].value
	v3 := rotated_items[(index+3000)%N].value

	fmt.Println(v1, v2, v3, v1+v2+v3)

}

func linked_list(data []int) {
	var node_list []*Node

	var node0 *Node

	for i, n := range data {
		node := Node{}
		node.value = n
		node_list = append(node_list, &node)
		if n == 0 {
			fmt.Println("ZERO", i)
			node0 = &node
		}
	}

	for i := 0; i < len(node_list); i++ {
		node_list[i].prev = node_list[(len(node_list)+i-1)%len(node_list)]
	}

	for i := 0; i < len(node_list); i++ {
		node_list[i].next = node_list[(i+1)%len(node_list)]
	}

	N := len(node_list)

	for mix_n := 0; mix_n < 10; mix_n++ {

		for _, node := range node_list {
			if node.value == 0 {
				continue
			}

			rotate := abs(node.value) % (N - 1)

			if node.value > 0 {
				P := node.prev
				for i := 0; i < rotate; i++ {
					X := node.next
					Y := X.next

					P.next = X
					X.prev = P
					X.next = node
					node.prev = X
					node.next = Y
					Y.prev = node

					P = X
				}
			}
			if node.value < 0 {
				P := node.next
				for i := 0; i < rotate; i++ {
					X := node.prev
					Y := X.prev

					P.prev = X
					X.next = P
					X.prev = node
					node.next = X
					node.prev = Y
					Y.next = node

					P = X
				}
			}
		}
	}
	node := node0
	res := 0
	for i := 0; i <= 3000; i++ {
		if i%1000 == 0 {
			fmt.Println(strNode(*node))
			res += node.value
		}
		node = node.next
	}
	fmt.Println(res)
}

func strNode(n Node) string {
	return fmt.Sprintf("(%d, %d, %d)", n.prev.value, n.value, n.next.value)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
