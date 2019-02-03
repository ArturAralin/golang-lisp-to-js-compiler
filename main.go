package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"parser"
	"path/filepath"
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

	file := absPath("./code")

	fmt.Println(file)

	d := readFile(file)

	fmt.Println(parser.Parse(file, d))
}
