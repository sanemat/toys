package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func buildCat() {
	if out, err := exec.Command("go", "build", "-o", filepath.Join("out", "cat"), "cat.go").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func buildEcho() {
	if out, err := exec.Command("go", "build", "-o", filepath.Join("out", "echo"), "echo.go").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func fmtAll() {
	if out, err := exec.Command("go", "fmt", "./...").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
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
