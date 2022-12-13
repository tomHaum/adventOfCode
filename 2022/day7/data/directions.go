package data

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

type File struct {
	Name string
	Size int
}
type Folder struct {
	Name       string
	SubFolders []*Folder
	Files      []*File
	Parent     *Folder
	Size       int
}

// GetData1 GetData true for up false for down
func GetData1(fileName string, maxSize int) (*Folder, []*Folder) {
	file, err := os.Open(fileName)
	Check(err)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	cdRegex := regexp.MustCompile(`\$ cd (.+)`)
	lsRegex := regexp.MustCompile(`\$ ls`)
	dirRegex := regexp.MustCompile(`dir (\w+)`)
	fileRegex := regexp.MustCompile(`(\d+) (.+)`)

	rootFolder := Folder{}
	folderUnderMaxSize := make([]*Folder, 0)
	currentFolder := &rootFolder
outerLoop:
	for scanner.Scan() {
		line := scanner.Text()

		matches := cdRegex.FindStringSubmatch(line)
		if len(matches) == 2 {
			name := matches[1]
			if name == "/" {
				folders := make([]*Folder, 0)
				files := make([]*File, 0)
				rootFolder = Folder{name, folders, files, &rootFolder, 0}
				continue
			}

			if name == ".." {
				calculateFolderSide(currentFolder)
				if currentFolder.Size < maxSize {
					folderUnderMaxSize = append(folderUnderMaxSize, currentFolder)
				}
				currentFolder = (*currentFolder).Parent
				continue
			}

			for _, f := range currentFolder.SubFolders {
				if f.Name == name {
					currentFolder = f
					continue outerLoop
				}
			}
			panic("impossible")
		}

		matches = lsRegex.FindStringSubmatch(line)
		if len(matches) == 1 {
			continue
		}

		matches = dirRegex.FindStringSubmatch(line)
		if len(matches) == 2 {
			name := matches[1]
			folders := make([]*Folder, 0)
			files := make([]*File, 0)
			folder := Folder{name, folders, files, currentFolder, 0}
			currentFolder.SubFolders = append(currentFolder.SubFolders, &folder)
			continue
		}

		matches = fileRegex.FindStringSubmatch(line)
		if len(matches) == 3 {
			size, err := strconv.Atoi(matches[1])
			Check(err)
			name := matches[2]
			file := File{name, size}
			currentFolder.Files = append(currentFolder.Files, &file)
			continue
		}
	}

	for currentFolder.Name != rootFolder.Name {
		calculateFolderSide(currentFolder)
		currentFolder = (*currentFolder).Parent
	}
	calculateFolderSide(&rootFolder)
	return &rootFolder, folderUnderMaxSize
}

func calculateFolderSide(currentFolder *Folder) {
	for _, f := range currentFolder.Files {
		currentFolder.Size += f.Size
	}
	for _, f := range currentFolder.SubFolders {
		currentFolder.Size += f.Size
	}
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
