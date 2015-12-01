// Exercise 1.4: Modify dup2 to print the names of all files in which each
// 				 duplicated line occurs.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, "stdin")
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, filename)
			f.Close()
		}
	}
}

func countLines(f *os.File, filename string) {
	input := bufio.NewScanner(f)
	counts := make(map[string]int)
	for input.Scan() {
		counts[input.Text()]++
	}
	flag := 0
	for lines, n := range counts {
		if n > 1 {
			flag = 1
			fmt.Printf("%d\t%s\n", n, lines)
		}
	}
	if flag == 1 {
		fmt.Println("The above duplicated lines are found in: ", filename)
	}
}
