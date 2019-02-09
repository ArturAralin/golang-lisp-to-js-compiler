package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"./src/jscompiler"
	"./src/parser"
)

func readFile(path string) string {
	d, _ := ioutil.ReadFile(path)

	return string(d)
}

func absPath(s string) string {
	dir, err := filepath.Abs(s)
	if err != nil {
		log.Fatal(err)
	}

	return dir
}

func main() {
	val := os.Args[1]

	ttree := parser.Parse("", val)

	fmt.Print(jscompiler.Compile(ttree))
}
