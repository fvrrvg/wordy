package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		Help()
		return
	}

	switch os.Args[1] {
	case "--common", "-c":
		Commonword()
	case "--rare", "-r":
		Rareword()
	case "--help", "-h":
		Help()
	case "--rules", "-R":
		Rules()
	default:
		fmt.Println("\033[31m" + "Invalid option" + "\x1b[0m")
		Help()
	}
}
