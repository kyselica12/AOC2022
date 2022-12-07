package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type folder struct {
	name       string
	parent     *folder
	subFolders map[string]folder
	files      map[string]int
}

func (f folder) addSubFolder(name string) {
	if _, ok := f.subFolders[name]; !ok {
		f.subFolders[name] = folder{parent: &f, name: name,
			subFolders: map[string]folder{},
			files:      map[string]int{}}
	}
}

func main() {
	input, _ := os.Open("./day7/input.txt")
	defer input.Close()
	sc := bufio.NewScanner(input)

	var root folder = folder{name: "root", subFolders: map[string]folder{}, files: map[string]int{}}
	root.addSubFolder("/")
	var current_folder folder = root
	names := map[string]bool{}

	for sc.Scan() {

		line := sc.Text()
		//fmt.Println(line)
		switch line[:4] {
		case "$ cd":
			name := line[5:]
			if name == ".." {
				current_folder = *current_folder.parent
			} else {
				//fmt.Println(name)
				names[name] = true
				current_folder = current_folder.subFolders[name]
			}
		case "$ ls":
			continue
		default:
			if line[:3] == "dir" {
				//fmt.Println(line[4:], current_folder.name, "AAAAAAAAAAAA")
				current_folder.addSubFolder(line[4:])
			} else {
				splitLine := strings.Split(line, " ")
				size, _ := strconv.Atoi(splitLine[0])
				name := splitLine[1]
				current_folder.files[name] = size
			}
		}

	}

	sizes := map[string]int{}
	size := root.getSize(sizes)
	fmt.Println(size)
	needed_size := 30000000 - (70000000 - size)

	diff := 70000000
	winner := "root"

	for name, size := range sizes {
		if size > needed_size && size-needed_size < diff {
			diff = size - needed_size
			winner = name
		}
	}

	fmt.Println(winner, sizes[winner])

}

func (f folder) getSize(sizes map[string]int) int {
	size := 0
	for _, value := range f.files {
		size += value
	}

	for _, subF := range f.subFolders {
		sub_size := subF.getSize(sizes)
		size += sub_size
	}

	//fmt.Println(f.name, size)
	sizes[f.name] = size

	return size
}
