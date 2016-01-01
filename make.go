package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func buildCommand(fileName string, commandName string) (string, error) {
	builtFile := filepath.Join("out", commandName)
	if _, err := exec.Command("go", "build", "-o", builtFile, fileName).Output(); err != nil {
		return builtFile, err
	}
	return builtFile, nil
}

func buildCat() {
	if out, err := buildCommand("cat.go", "cat"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func buildEcho() {
	if out, err := buildCommand("echo.go", "echo"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func buildSieve() {
	if out, err := buildCommand("sieve.go", "sieve"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func buildHead() {
	if out, err := buildCommand("head.go", "head"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func buildTail() {
	if out, err := buildCommand("tail.go", "tail"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func fmtAll() {
	if out, err := exec.Command("go", "fmt", "./...").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func beforeScript() {
	if out, err := exec.Command("go", "get", "golang.org/x/tools/cmd/vet").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
	if out, err := exec.Command("go", "get", "golang.org/x/tools/cmd/goimports").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
	if out, err := exec.Command("go", "get", "github.com/golang/lint/golint").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goVet() {
	if out, err := exec.Command("go", "vet", "./...").Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goLint() {
	if out, err := exec.Command("golint", "./...").Output(); err != nil {
		log.Fatal(err)
	} else if len(out) != 0 {
		log.Fatal(string(out))
	}
}

func goImports() {
	if out, err := exec.Command("goimports", "-l", "./").Output(); err != nil {
		log.Fatal(err)
	} else if len(out) != 0 {
		log.Fatal(string(out))
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
		buildSieve()
		buildHead()
		buildTail()
	case "fmt":
		fmtAll()
	case "beforescript":
		beforeScript()
	case "govet":
		goVet()
	case "golint":
		goLint()
	case "goimports":
		goImports()
	case "script":
		goVet()
		goLint()
		goImports()
	default:
		panic("Unknown action: " + action)
	}
}
