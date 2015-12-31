package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
)

func buildCat() {
	out, err := exec.Command("go", "build", "-o", filepath.Join("out", "cat"), "cat.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}

func buildEcho() {
	out, err := exec.Command("go", "build", "-o", filepath.Join("out", "echo"), "echo.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}

func main() {
	buildCat()
	buildEcho()
}
