package main

import (
	"adventofcode2023/day05"
	"log"
	"os"
	"runtime/pprof"
	"time"
)
import _ "embed"

//go:embed part1input
var input string

func main() {
	f, err := os.Create("cpu-profile")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	go func() {
		time.Sleep(time.Minute)
		pprof.StopCPUProfile()
	}()
	_, _ = day05.Run(input)
}
