package main

import (
	"adventOfCode2022/day7/data"
	"fmt"
)

const totalDiskSpace = 70000000
const updateSize = 30000000

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

func main() {
	fmt.Println("starting day 2")
	fileName := "data/input.txt"
	root, smallishFolders := data.GetData1(fileName, 100000)
	sum := 0
	for _, f := range smallishFolders {
		sum += f.Size
	}
	spaceRemaining := totalDiskSpace - root.Size
	searchFor := updateSize - spaceRemaining
	foldersToDelete := make([]*data.Folder, 0)
	foldersToDelete = walkTree(root, searchFor, foldersToDelete)
	smallest := foldersToDelete[0].Size
	for _, f := range foldersToDelete {
		if smallest > f.Size {
			smallest = f.Size
		}
	}
	fmt.Printf("part 1: %v\n", sum)
	fmt.Printf("part 2: %v\n", smallest)
}

func walkTree(root *data.Folder, search int, foldersToDelete []*data.Folder) []*data.Folder {
	if root.Size > search {
		foldersToDelete = append(foldersToDelete, root)
	}
	for _, f := range root.SubFolders {
		foldersToDelete = walkTree(f, search, foldersToDelete)
	}
	return foldersToDelete
}
