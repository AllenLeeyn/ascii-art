package main

import (
	"ascii-art/pkg/generator"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	argsLen := len(args)

	if argsLen != 1 {
		fmt.Println("Usage: go run . <something>")
		fmt.Println("EX: go run . \"Hello There\"")
		return
	}

	txt := args[0]
	styleNm := "standard"
	fmt.Print(generator.GenArt(txt, styleNm))
}
