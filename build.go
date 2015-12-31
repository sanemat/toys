package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func main() {
	out, err := exec.Command("go", "build", "-o", filepath.Join("out", "cat"), "cat.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
