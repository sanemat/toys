package main

import (
	"fmt"
	"log"
	"os"
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

func fmtAll() {
	out, err := exec.Command("go", "fmt", "./...").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}

func main() {
	action := ""
	if len(os.Args) == 2 {
		action = os.Args[1]
	}

	switch action {
	case "build":
		buildCat()
		buildEcho()
	case "fmt":
		fmtAll()
	default:
		panic("Unknown action: " + action)
	}
}
