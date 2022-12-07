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
	input, _ := os.Open("./day7/test-input.txt")
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

	size, good_size := root.getSize()
	fmt.Println(size, good_size)
}

func (f folder) getSize() (int, int) {
	size := 0
	good_size := 0
	for _, value := range f.files {
		size += value
	}

	for _, subF := range f.subFolders {
		sub_size, sub_good_size := subF.getSize()
		size += sub_size
		good_size += sub_good_size
		if sub_size <= 100000 {
			good_size += sub_size
		}
	}

	fmt.Println(f.name, size, good_size)

	return size, good_size
}
