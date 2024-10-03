package main

import (
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Requires at least one arg!")
	}

	var args []string
	if len(os.Args) > 2 {
		args = os.Args[2:]
	}

	switch cmd := os.Args[1]; cmd {
	case "info":
		cmdInfo(args)
	case "styled":
		cmdStyled(args)
	default:
		log.Fatalf("Invalid command: %q", cmd)
	}
}
