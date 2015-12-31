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
	for scanner.Scan() {
		fmt.Println(scanner.Text()) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
