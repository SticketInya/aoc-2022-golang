package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("AOC 2022 - Day 7")

	task_1()
}

type directory struct {
	Name     string
	FileSize int
	IsFile   bool
	Children map[string]*directory
	Parent   *directory
}

func task_1() {
	file, err := os.Open("day_07/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var current *directory
	fileSystem := directory{}

	for scanner.Scan() {
		line := strings.Fields(scanner.Text())
		if line[1] == "ls" {
			continue
		}

		if len(line) > 2 {
			commandTarget := line[2]

			if commandTarget == ".." {
				current = current.Parent
				continue
			}

			if commandTarget == "/" {
				fileSystem = directory{
					Name:     commandTarget,
					Children: make(map[string]*directory),
				}
				current = &fileSystem
				continue
			}

			current = current.Children[commandTarget]
			continue
		}

		if line[0] == "dir" {
			dirName := line[1]

			current.Children[dirName] = &directory{
				Name:     dirName,
				Parent:   current,
				Children: make(map[string]*directory),
			}
			continue
		}

		if line[0] != "$" {
			fileName := line[1]

			size, _ := strconv.Atoi(line[0])
			current.Children[fileName] = &directory{
				Name:     fileName,
				FileSize: size,
				IsFile:   true,
				Parent:   current,
			}
		}

	}

	// for _, child := range fileSystem.Children {
	// 	fmt.Printf("%+v\n", child)

	// }
	cache := stringCache{
		cache: map[string]int{},
	}
	cache.calculateSize(fileSystem)

	for key, val := range cache.cache {
		fmt.Printf("dir %s size: %v\n", key, val)
	}

	fmt.Printf("The result: %v\n", calculateTotal(cache.cache, 100000))

}

func calculateTotal(values map[string]int, limit int) int {
	total := 0
	for _, val := range values {
		if val <= limit {
			total += val
		}
	}
	return total
}

type stringCache struct {
	cache map[string]int
}

func (s *stringCache) calculateSize(root directory) int {
	size := 0
	if root.IsFile {
		return root.FileSize
	}

	cachedVal, ok := s.cache[root.Name]
	if ok {
		return cachedVal
	}

	for _, childDir := range root.Children {
		size += s.calculateSize(*childDir)
	}
	s.cache[root.Name] = size
	return size
}
