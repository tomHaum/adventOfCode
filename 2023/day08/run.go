package day08

import (
	"fmt"
	"strings"
	"sync"
)

func run(input string) (int, int) {
	input = strings.ReplaceAll(input, "\r\n", "\n")

	split := strings.Split(input, "\n")
	instructions := split[0]

	split = split[2:]
	nodes := make(map[string]*Node)
	for _, l := range split {
		n := ParseNode(l)
		nodes[n.Name] = &n
	}

	for _, n := range nodes {
		n.LeftNode = nodes[n.Left]
		n.RightNode = nodes[n.Right]
	}

	count := 0
	//currNode := nodes["AAA"]
	//for {
	//	count++
	//	if instructions[i] == 'L' {
	//		currNode = currNode.LeftNode
	//	} else {
	//		currNode = currNode.RightNode
	//	}
	//
	//	if currNode.Name == "ZZZ" {
	//		break
	//	}
	//	i++
	//	if i >= len(instructions) {
	//		i = 0
	//	}
	//}
	part1 := count

	ghostNodes := make([]*Node, 0)
	for _, n := range nodes {
		if strings.HasSuffix(n.Name, "A") {
			ghostNodes = append(ghostNodes, n)
		}
	}

	doneWg := sync.WaitGroup{}
	doneWg.Add(1)
	wg := sync.WaitGroup{}
	wg.Add(len(ghostNodes))

	mutext := sync.Mutex{}

	maxGhostCount := 0
	ghostCounts := make([]int, len(ghostNodes))
	findGhostNode := func(ghostIndex int) {
		currNode := ghostNodes[ghostIndex]
		visited := make(map[string]interface{})
		visited[currNode.Name] = struct{}{}

		i := 0
		count := 0
		repeats := 0
		for {
			count++
			if instructions[i] == 'L' {
				currNode = currNode.LeftNode
			} else {
				currNode = currNode.RightNode
			}
			if _, ok := visited[currNode.Name]; ok {
				//fmt.Printf("vist again found: %s %s %v\n", ghostNodes[0].Name, currNode.Name, count)
			}
			visited[currNode.Name] = struct{}{}

			if currNode.Name[2] == 'Z' {
				fmt.Printf("z found: %s %s %v\n", ghostNodes[0].Name, currNode.Name, count)
				ghostCounts[ghostIndex] = count
				wg.Done()
				for {
					wg.Wait()

					mutext.Lock()
					equalCount := 0
					for _, mc := range ghostCounts {
						if mc > maxGhostCount {
							maxGhostCount = mc
						}
						if mc == maxGhostCount {
							equalCount++
						}
					}

					if equalCount == len(ghostCounts) {
						doneWg.Done()
						return
					}
					mutext.Unlock()

					if count < maxGhostCount {
						break
					} else {
						continue // wait until someone else gets max
					}
				}

				break
			}
			i++
			if i >= len(instructions) {
				i = 0
				repeats += 1
			}
		}

	}
	for i, _ := range ghostNodes {
		go findGhostNode(i)
	}
	doneWg.Wait()

	return part1, maxGhostCount
}

type Node struct {
	Name string

	Left      string
	LeftNode  *Node
	Right     string
	RightNode *Node
}

func ParseNode(line string) Node {
	name, rest, _ := strings.Cut(line, "=")

	rest = strings.ReplaceAll(rest, "(", "")
	rest = strings.ReplaceAll(rest, ")", "")
	rest = strings.ReplaceAll(rest, " ", "")
	rest = strings.TrimSpace(rest)
	name = strings.TrimSpace(name)

	left, right, _ := strings.Cut(rest, ",")

	n := Node{
		Name:  name,
		Left:  left,
		Right: right,
	}
	return n
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
