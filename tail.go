package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// https://golang.org/pkg/bufio/#example_Scanner_lines
	// http://qiita.com/ikawaha/items/28186d965780fab5533d
	var fp *os.File
	var err error
	const limit = 10

	flag.Parse()

	if flag.NArg() == 0 {
		fp = os.Stdin
	} else {
		fp, err = os.Open(flag.Arg(0))
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()
	}
	scanner := bufio.NewScanner(fp)
	capturedLines := []string{}
	for scanner.Scan() {
		capturedLines = append(capturedLines, scanner.Text())
	}

	lineLength := len(capturedLines)
	reversedLines := []string{}
	for i := 0; i < limit && i < lineLength; i++ {
		reversedLines = append(reversedLines, capturedLines[lineLength-1-i])
	}
	reversedLineLength := len(reversedLines)
	for i := 0; i < reversedLineLength; i++ {
		fmt.Println(reversedLines[reversedLineLength-1-i])
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
