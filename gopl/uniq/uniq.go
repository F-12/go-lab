package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("no file path is provided")
		os.Exit(1)
	}
	for _, p := range args {
		handleFile(p)
	}
}

func handleFile(path string) {
	if f, err := os.Open(path); err == nil {
		counter := make(map[string]int)
		countLines(f, counter)
		// printCounter(counter)
	}
}
func countLines(f *os.File, counter map[string]int) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		counter[line]++
		if counter[line] > 1 {
			fmt.Printf("encounter duplicate lines in %s\n", f.Name())
		}
	}
}

func printCounter(counter map[string]int) {
	for k, v := range counter {
		fmt.Printf("%s : %v\n", k, v)
	}
}
