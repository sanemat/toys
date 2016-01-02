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
	bytes := 0
	for scanner.Scan() {
		bytes += len(scanner.Bytes())
	}
	fmt.Println(bytes)
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
