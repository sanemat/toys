package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("go", "build", "-o", "out/cat", "cat.go").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(out))
}
