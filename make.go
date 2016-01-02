package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func buildCommand(fileName string, commandName string) (string, error) {
	var commandNameWithExt string
	if runtime.GOOS == "windows" {
		commandNameWithExt = commandName + ".exe"
	} else {
		commandNameWithExt = commandName
	}
	builtFile := filepath.Join("out", commandNameWithExt)
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

func buildCountByte() {
	if out, err := buildCommand("countbyte.go", "countbyte"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(out)
	}
}

func fmtAll() {
	comm := []string{"go", "fmt", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func beforeScript() {
	var comm []string
	comm = []string{"go", "get", "golang.org/x/tools/cmd/vet"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
	comm = []string{"go", "get", "golang.org/x/tools/cmd/goimports"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
	comm = []string{"go", "get", "github.com/golang/lint/golint"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goVet() {
	comm := []string{"go", "vet", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else {
		fmt.Print(string(out))
	}
}

func goLint() {
	comm := []string{"golint", "./..."}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
		log.Fatal(err)
	} else if len(out) != 0 {
		log.Fatal(string(out))
	}
}

func goImports() {
	comm := []string{"goimports", "-l", "./"}
	if out, err := exec.Command(comm[0], comm[1:]...).Output(); err != nil {
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
		buildCountByte()
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
