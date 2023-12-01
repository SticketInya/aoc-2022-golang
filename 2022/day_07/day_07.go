package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("AOC 2022 - Day 7")

	task_1()
}

type directory struct {
	Id       uuid.UUID
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
					Id:       uuid.New(),
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
				Id:       uuid.New(),
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
				Id:       uuid.New(),
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
		cache: map[uuid.UUID]int{},
	}
	cache.calculateSize(fileSystem)

	// for key, val := range cache.cache {
	// 	fmt.Printf("dir %s size: %v\n", key, val)
	// }

	fmt.Printf("The result for task 1: %v\n", calculateTotal(cache.cache, 100000))

	sizeToDelete := getSizeToDelete(fileSystem, cache.cache)

	fmt.Printf("The result for task 2: %v\n", getTheSmallestDirectoryToDelete(sizeToDelete, cache.cache))

}

func calculateTotal(values map[uuid.UUID]int, limit int) int {
	total := 0
	for _, val := range values {
		if val <= limit {
			total += val
		}
	}
	return total
}

func getSizeToDelete(root directory, sizesCache map[uuid.UUID]int) int {
	const TOTAL_DISK_SPACE = 70000000
	const MIN_FREE_SPACE = 30000000
	size, ok := sizesCache[root.Id]
	if !ok {
		return -1
	}

	return MIN_FREE_SPACE - (TOTAL_DISK_SPACE - size)
}

func getTheSmallestDirectoryToDelete(sizeToDelete int, dirCache map[uuid.UUID]int) int {
	possibleSizes := []int{}

	for _, size := range dirCache {
		if size >= sizeToDelete {
			possibleSizes = append(possibleSizes, size)
		}
	}

	sort.Ints(possibleSizes)

	return possibleSizes[0]
}

type stringCache struct {
	cache map[uuid.UUID]int
}

func (s *stringCache) calculateSize(root directory) int {
	size := 0
	if root.IsFile {
		return root.FileSize
	}

	for _, childDir := range root.Children {
		size += s.calculateSize(*childDir)
	}
	s.cache[root.Id] = size
	return size
}
