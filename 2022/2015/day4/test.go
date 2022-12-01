package main

import (
	"crypto/md5"
	"fmt"
	"time"
)

func main() {
	fmt.Println("starting day 4")
	startTime := time.Now()
	key := "bgvyzdsv"
	i := 0
	found1 := false
	found2 := false

	for {
		hashKey := fmt.Sprint(key, i)
		result := md5.Sum([]byte(hashKey))
		if result[0] == 0 && result[1] == 0 {
			if !found1 && result[2]&0b11110000 == 0 {
				fmt.Printf("Part 1 | Iteration: %v\n", i)
				found1 = true
			}
			if result[2] == 0 {
				fmt.Printf("Part 2 | Iteration: %v\n", i)
				found2 = true
			}
			if found1 && found2 {
				break
			}
		}
		i++
	}
	endTime := time.Now()

	fmt.Printf("Runtime: %v", endTime.Sub(startTime).Nanoseconds())
}
